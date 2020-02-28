package search

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"go.skia.org/infra/go/paramtools"
	"go.skia.org/infra/go/testutils"
	"go.skia.org/infra/go/testutils/unittest"
	"go.skia.org/infra/go/tiling"
	mock_clstore "go.skia.org/infra/golden/go/clstore/mocks"
	"go.skia.org/infra/golden/go/code_review"
	"go.skia.org/infra/golden/go/comment"
	mock_comment "go.skia.org/infra/golden/go/comment/mocks"
	"go.skia.org/infra/golden/go/comment/trace"
	"go.skia.org/infra/golden/go/diff"
	mock_diffstore "go.skia.org/infra/golden/go/diffstore/mocks"
	"go.skia.org/infra/golden/go/digest_counter"
	"go.skia.org/infra/golden/go/expectations"
	mock_expectations "go.skia.org/infra/golden/go/expectations/mocks"
	"go.skia.org/infra/golden/go/indexer"
	mock_index "go.skia.org/infra/golden/go/indexer/mocks"
	"go.skia.org/infra/golden/go/paramsets"
	"go.skia.org/infra/golden/go/search/common"
	"go.skia.org/infra/golden/go/search/frontend"
	"go.skia.org/infra/golden/go/search/query"
	data "go.skia.org/infra/golden/go/testutils/data_three_devices"
	"go.skia.org/infra/golden/go/tjstore"
	mock_tjstore "go.skia.org/infra/golden/go/tjstore/mocks"
	"go.skia.org/infra/golden/go/types"
)

// TODO(kjlubick) Add tests for:
//   - When a CL doesn't exist or the CL has not patchsets, patchset doesn't exist,
//     or otherwise no results.
//   - Use ignore matcher
//   - When a CL specifies a PS
//   - IncludeMaster=true
//   - UnavailableDigests is not empty
//   - DiffSever/RefDiffer error

// TestSearchThreeDevicesSunnyDay searches over the three_devices
// test data for untriaged images at head, essentially the default search.
// We expect to get two untriaged digests, with their closest positive and
// negative images (if any).
func TestSearchThreeDevicesSunnyDay(t *testing.T) {
	unittest.SmallTest(t)

	mds := makeDiffStoreWithNoFailures()
	addDiffData(mds, data.AlphaUntriaged1Digest, data.AlphaGood1Digest, makeSmallDiffMetric())
	addDiffData(mds, data.AlphaUntriaged1Digest, data.AlphaBad1Digest, makeBigDiffMetric())
	addDiffData(mds, data.BetaUntriaged1Digest, data.BetaGood1Digest, makeBigDiffMetric())
	// BetaUntriaged1Digest has no negative images to compare against.

	s := New(mds, makeThreeDevicesExpectationStore(), makeThreeDevicesIndexer(), nil, nil, emptyCommentStore(), everythingPublic)

	q := &query.Search{
		ChangeListID: "",
		Unt:          true,
		Head:         true,

		Metric:   diff.CombinedMetric,
		FRGBAMin: 0,
		FRGBAMax: 255,
		FDiffMax: -1,
		Sort:     query.SortAscending,
	}

	resp, err := s.Search(context.Background(), q)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, &frontend.SearchResponse{
		Commits: data.MakeTestCommits(),
		Offset:  0,
		Size:    2,
		Digests: []*frontend.SRDigest{
			// AlphaTest comes first because we are sorting by ascending
			// "combined" metric, and AlphaTest's closest match is the
			// small diff metric, whereas BetaTest's only match is the
			// big diff metric.
			{
				Test:   data.AlphaTest,
				Digest: data.AlphaUntriaged1Digest,
				Status: "untriaged",
				ParamSet: map[string][]string{
					"device":              {data.BullheadDevice},
					types.PrimaryKeyField: {string(data.AlphaTest)},
					types.CorpusField:     {"gm"},
				},
				Traces: &frontend.TraceGroup{
					TileSize: 3, // 3 commits in tile
					Traces: []frontend.Trace{
						{
							Data: []int{1, 1, 0},
							ID:   data.BullheadAlphaTraceID,
							Params: map[string]string{
								"device":              data.BullheadDevice,
								types.PrimaryKeyField: string(data.AlphaTest),
								types.CorpusField:     "gm",
							},
						},
					},
					Digests: []frontend.DigestStatus{
						{
							Digest: data.AlphaUntriaged1Digest,
							Status: "untriaged",
						},
						{
							Digest: data.AlphaBad1Digest,
							Status: "negative",
						},
					},
				},
				ClosestRef: common.PositiveRef,
				RefDiffs: map[common.RefClosest]*frontend.SRDiffDigest{
					common.PositiveRef: {
						DiffMetrics: makeSmallDiffMetric(),
						Digest:      data.AlphaGood1Digest,
						Status:      "positive",
						ParamSet: map[string][]string{
							"device":              {data.AnglerDevice, data.CrosshatchDevice},
							types.PrimaryKeyField: {string(data.AlphaTest)},
							types.CorpusField:     {"gm"},
						},
						OccurrencesInTile: 2,
					},
					common.NegativeRef: {
						DiffMetrics: makeBigDiffMetric(),
						Digest:      data.AlphaBad1Digest,
						Status:      "negative",
						ParamSet: map[string][]string{
							"device":              {data.AnglerDevice, data.BullheadDevice, data.CrosshatchDevice},
							types.PrimaryKeyField: {string(data.AlphaTest)},
							types.CorpusField:     {"gm"},
						},
						OccurrencesInTile: 6,
					},
				},
			},
			{
				Test:   data.BetaTest,
				Digest: data.BetaUntriaged1Digest,
				Status: "untriaged",
				ParamSet: map[string][]string{
					"device":              {data.CrosshatchDevice},
					types.PrimaryKeyField: {string(data.BetaTest)},
					types.CorpusField:     {"gm"},
				},
				Traces: &frontend.TraceGroup{
					TileSize: 3,
					Traces: []frontend.Trace{
						{
							Data: []int{0, missingDigestIndex, missingDigestIndex},
							ID:   data.CrosshatchBetaTraceID,
							Params: map[string]string{
								"device":              data.CrosshatchDevice,
								types.PrimaryKeyField: string(data.BetaTest),
								types.CorpusField:     "gm",
							},
						},
					},
					Digests: []frontend.DigestStatus{
						{
							Digest: data.BetaUntriaged1Digest,
							Status: "untriaged",
						},
					},
				},
				ClosestRef: common.PositiveRef,
				RefDiffs: map[common.RefClosest]*frontend.SRDiffDigest{
					common.PositiveRef: {
						DiffMetrics: makeBigDiffMetric(),
						Digest:      data.BetaGood1Digest,
						Status:      "positive",
						ParamSet: map[string][]string{
							"device":              {data.AnglerDevice, data.BullheadDevice},
							types.PrimaryKeyField: {string(data.BetaTest)},
							types.CorpusField:     {"gm"},
						},
						OccurrencesInTile: 6,
					},
					common.NegativeRef: nil,
				},
			},
		},
	}, resp)
}

// TestSearchThreeDevicesQueries searches over the three_devices test data using a variety
// of queries. It only spot-checks the returned data (e.g. things are in the right order); other
// tests should do a more thorough check of the return values.
func TestSearchThreeDevicesQueries(t *testing.T) {
	unittest.SmallTest(t)

	mds := makeDiffStoreWithNoFailures()
	addDiffData(mds, data.AlphaUntriaged1Digest, data.AlphaGood1Digest, makeSmallDiffMetric())
	addDiffData(mds, data.AlphaUntriaged1Digest, data.AlphaBad1Digest, makeBigDiffMetric())
	addDiffData(mds, data.BetaUntriaged1Digest, data.BetaGood1Digest, makeBigDiffMetric())
	// BetaUntriaged1Digest has no negative images to compare against.

	s := New(mds, makeThreeDevicesExpectationStore(), makeThreeDevicesIndexer(), nil, nil, emptyCommentStore(), everythingPublic)

	// spotCheck is the subset of data we assert against.
	type spotCheck struct {
		test            types.TestName
		digest          types.Digest
		labelStr        string
		closestPositive types.Digest
		closestNegative types.Digest
	}

	test := func(name string, input *query.Search, expectedOutputs []spotCheck) {
		t.Run(name, func(t *testing.T) {
			resp, err := s.Search(context.Background(), input)
			require.NoError(t, err)
			require.NotNil(t, resp)

			require.Len(t, resp.Digests, len(expectedOutputs))
			for i, actualDigest := range resp.Digests {
				expected := expectedOutputs[i]
				assert.Equal(t, expected.test, actualDigest.Test)
				assert.Equal(t, expected.digest, actualDigest.Digest)
				assert.Equal(t, expected.labelStr, actualDigest.Status)
				if expected.closestPositive == "" {
					assert.Nil(t, actualDigest.RefDiffs[common.PositiveRef])
				} else {
					cp := actualDigest.RefDiffs[common.PositiveRef]
					require.NotNil(t, cp)
					assert.Equal(t, expected.closestPositive, cp.Digest)
				}
				if expected.closestNegative == "" {
					assert.Nil(t, actualDigest.RefDiffs[common.NegativeRef])
				} else {
					cp := actualDigest.RefDiffs[common.NegativeRef]
					require.NotNil(t, cp)
					assert.Equal(t, expected.closestNegative, cp.Digest)
				}
			}
		})
	}

	test("default query, but in reverse", &query.Search{
		Unt:  true,
		Head: true,

		Metric:   diff.CombinedMetric,
		FRGBAMin: 0,
		FRGBAMax: 255,
		FDiffMax: -1,
		Sort:     query.SortDescending,
	}, []spotCheck{
		{
			test:            data.BetaTest,
			digest:          data.BetaUntriaged1Digest,
			labelStr:        "untriaged",
			closestPositive: data.BetaGood1Digest,
			closestNegative: "",
		},
		{
			test:            data.AlphaTest,
			digest:          data.AlphaUntriaged1Digest,
			labelStr:        "untriaged",
			closestPositive: data.AlphaGood1Digest,
			closestNegative: data.AlphaBad1Digest,
		},
	})

	test("the closest RGBA diff should be at least 50 units away", &query.Search{
		Unt:  true,
		Head: true,

		Metric:   diff.CombinedMetric,
		FRGBAMin: 50,
		FRGBAMax: 255,
		FDiffMax: -1,
		Sort:     query.SortDescending,
	}, []spotCheck{
		{
			test:            data.BetaTest,
			digest:          data.BetaUntriaged1Digest,
			labelStr:        "untriaged",
			closestPositive: data.BetaGood1Digest,
			closestNegative: "",
		},
	})

	// note: this matches only the makeSmallDiffMetric
	test("the closest RGBA diff should be no more than 50 units away", &query.Search{
		Unt:  true,
		Head: true,

		Metric:   diff.CombinedMetric,
		FRGBAMin: 0,
		FRGBAMax: 50,
		FDiffMax: -1,
		Sort:     query.SortDescending,
	}, []spotCheck{
		{
			test:            data.AlphaTest,
			digest:          data.AlphaUntriaged1Digest,
			labelStr:        "untriaged",
			closestPositive: data.AlphaGood1Digest,
			closestNegative: data.AlphaBad1Digest,
		},
	})

	test("combined diff metric less than 1", &query.Search{
		Unt:  true,
		Head: true,

		Metric:   diff.CombinedMetric,
		FRGBAMin: 0,
		FRGBAMax: 255,
		FDiffMax: 1,
		Sort:     query.SortDescending,
	}, []spotCheck{
		{
			test:            data.AlphaTest,
			digest:          data.AlphaUntriaged1Digest,
			labelStr:        "untriaged",
			closestPositive: data.AlphaGood1Digest,
			closestNegative: data.AlphaBad1Digest,
		},
	})

	test("percent diff metric less than 1", &query.Search{
		Unt:  true,
		Head: true,

		Metric:   diff.PercentMetric,
		FRGBAMin: 0,
		FRGBAMax: 255,
		FDiffMax: 1,
		Sort:     query.SortDescending,
	}, []spotCheck{
		{
			test:            data.AlphaTest,
			digest:          data.AlphaUntriaged1Digest,
			labelStr:        "untriaged",
			closestPositive: data.AlphaGood1Digest,
			closestNegative: data.AlphaBad1Digest,
		},
	})

	test("Fewer than 10 different pixels", &query.Search{
		Unt:  true,
		Head: true,

		Metric:   diff.PixelMetric,
		FRGBAMin: 0,
		FRGBAMax: 255,
		FDiffMax: 10,
		Sort:     query.SortDescending,
	}, []spotCheck{
		{
			test:            data.AlphaTest,
			digest:          data.AlphaUntriaged1Digest,
			labelStr:        "untriaged",
			closestPositive: data.AlphaGood1Digest,
			closestNegative: data.AlphaBad1Digest,
		},
	})

	test("Nothing has fewer than 10 different pixels and min RGBA diff >50", &query.Search{
		Unt:  true,
		Head: true,

		Metric:   diff.PixelMetric,
		FRGBAMin: 50,
		FRGBAMax: 255,
		FDiffMax: 10,
		Sort:     query.SortDescending,
	}, nil)

	test("default query, only those with a reference diff (all of them)", &query.Search{
		Unt:  true,
		Head: true,
		FRef: true,

		Metric:   diff.CombinedMetric,
		FRGBAMin: 0,
		FRGBAMax: 255,
		FDiffMax: -1,
		Sort:     query.SortAscending,
	}, []spotCheck{
		{
			test:            data.AlphaTest,
			digest:          data.AlphaUntriaged1Digest,
			labelStr:        "untriaged",
			closestPositive: data.AlphaGood1Digest,
			closestNegative: data.AlphaBad1Digest,
		},
		{
			test:            data.BetaTest,
			digest:          data.BetaUntriaged1Digest,
			labelStr:        "untriaged",
			closestPositive: data.BetaGood1Digest,
			closestNegative: "",
		},
	})

	test("starting at the second commit, we only see alpha's untriaged commit at head", &query.Search{
		Unt:          true,
		Head:         true,
		FCommitBegin: data.MakeTestCommits()[1].Hash,

		Metric:   diff.CombinedMetric,
		FRGBAMin: 0,
		FRGBAMax: 255,
		FDiffMax: -1,
		Sort:     query.SortAscending,
	}, []spotCheck{
		{
			test:            data.AlphaTest,
			digest:          data.AlphaUntriaged1Digest,
			labelStr:        "untriaged",
			closestPositive: data.AlphaGood1Digest,
			closestNegative: data.AlphaBad1Digest,
		},
	})

	test("starting at the second commit, we see both if we ignore the head restriction", &query.Search{
		Unt:          true,
		Head:         false,
		FCommitBegin: data.MakeTestCommits()[1].Hash,

		Metric:   diff.CombinedMetric,
		FRGBAMin: 0,
		FRGBAMax: 255,
		FDiffMax: -1,
		Sort:     query.SortAscending,
	}, []spotCheck{
		{
			test:            data.AlphaTest,
			digest:          data.AlphaUntriaged1Digest,
			labelStr:        "untriaged",
			closestPositive: data.AlphaGood1Digest,
			closestNegative: data.AlphaBad1Digest,
		},
		{
			test:            data.BetaTest,
			digest:          data.BetaUntriaged1Digest,
			labelStr:        "untriaged",
			closestPositive: data.BetaGood1Digest,
			closestNegative: "",
		},
	})

	test("stopping at the second commit, we only see beta's untriaged", &query.Search{
		Unt:        true,
		Head:       true,
		FCommitEnd: data.MakeTestCommits()[1].Hash,

		Metric:   diff.CombinedMetric,
		FRGBAMin: 0,
		FRGBAMax: 255,
		FDiffMax: -1,
		Sort:     query.SortAscending,
	}, []spotCheck{
		{
			test:            data.BetaTest,
			digest:          data.BetaUntriaged1Digest,
			labelStr:        "untriaged",
			closestPositive: data.BetaGood1Digest,
			closestNegative: "",
		},
	})

	test("query matches nothing", &query.Search{
		Unt:  true,
		Head: true,
		TraceValues: map[string][]string{
			"blubber": {"nothing"},
		},

		Metric:   diff.CombinedMetric,
		FRGBAMin: 0,
		FRGBAMax: 255,
		FDiffMax: -1,
		Sort:     query.SortDescending,
	}, []spotCheck{})
}

// TestSearch_ThreeDevicesCorpusWithComments_CommentsInResults ensures that search results contain
// comments when it matches the traces.
func TestSearch_ThreeDevicesCorpusWithComments_CommentsInResults(t *testing.T) {
	unittest.SmallTest(t)

	bullheadComment := trace.Comment{
		ID:        "1",
		CreatedBy: "zulu@example.com",
		UpdatedBy: "zulu@example.com",
		CreatedTS: time.Date(2020, time.February, 19, 18, 17, 16, 0, time.UTC),
		UpdatedTS: time.Date(2020, time.February, 19, 18, 17, 16, 0, time.UTC),
		Comment:   "All bullhead devices draw upside down",
		QueryToMatch: paramtools.ParamSet{
			"device": []string{data.BullheadDevice},
		},
	}

	alphaTestComment := trace.Comment{
		ID:        "2",
		CreatedBy: "yankee@example.com",
		UpdatedBy: "xray@example.com",
		CreatedTS: time.Date(2020, time.February, 2, 18, 17, 16, 0, time.UTC),
		UpdatedTS: time.Date(2020, time.February, 20, 18, 17, 16, 0, time.UTC),
		Comment:   "Watch pixel 0,4 to make sure it's not purple",
		QueryToMatch: paramtools.ParamSet{
			types.PrimaryKeyField: []string{string(data.AlphaTest)},
		},
	}

	betaTestBullheadComment := trace.Comment{
		ID:        "4",
		CreatedBy: "victor@example.com",
		UpdatedBy: "victor@example.com",
		CreatedTS: time.Date(2020, time.February, 22, 18, 17, 16, 0, time.UTC),
		UpdatedTS: time.Date(2020, time.February, 22, 18, 17, 16, 0, time.UTC),
		Comment:   "Being upside down, this test should be ABGR instead of RGBA",
		QueryToMatch: paramtools.ParamSet{
			"device":              []string{data.BullheadDevice},
			types.PrimaryKeyField: []string{string(data.BetaTest)},
		},
	}

	commentAppliesToNothing := trace.Comment{
		ID:        "3",
		CreatedBy: "uniform@example.com",
		UpdatedBy: "uniform@example.com",
		CreatedTS: time.Date(2020, time.February, 26, 26, 26, 26, 0, time.UTC),
		UpdatedTS: time.Date(2020, time.February, 26, 26, 26, 26, 0, time.UTC),
		Comment:   "On Wednesdays, this device draws pink",
		QueryToMatch: paramtools.ParamSet{
			"device": []string{"This device does not exist"},
		},
	}

	mcs := &mock_comment.Store{}
	// Return these in an arbitrary, unsorted order
	mcs.On("ListComments", testutils.AnyContext).Return([]trace.Comment{commentAppliesToNothing, alphaTestComment, betaTestBullheadComment, bullheadComment}, nil)

	s := New(makeStubDiffStore(), makeThreeDevicesExpectationStore(), makeThreeDevicesIndexer(), nil, nil, mcs, everythingPublic)

	q := &query.Search{
		// Set all to true so all 6 traces show up in the final results.
		Unt:  true,
		Pos:  true,
		Neg:  true,
		Head: true,

		Metric:   diff.CombinedMetric,
		FRGBAMin: 0,
		FRGBAMax: 255,
		FDiffMax: -1,
		Sort:     query.SortAscending,
	}

	resp, err := s.Search(context.Background(), q)
	require.NoError(t, err)
	require.NotNil(t, resp)
	// There are 4 unique digests at HEAD on the three_devices corpus. Do a quick smoke test to make
	// sure we have one search result for each of them.
	require.Len(t, resp.Digests, 4)
	f := frontend.ToTraceComment
	// This should be sorted by UpdatedTS.
	assert.Equal(t, []frontend.TraceComment{
		f(bullheadComment), f(alphaTestComment), f(betaTestBullheadComment), f(commentAppliesToNothing),
	}, resp.TraceComments)

	// These numbers are indices into the resp.TraceComments. The nil entries are expected to have
	// no comments that match them.
	expectedComments := map[tiling.TraceID][]int{
		data.AnglerAlphaTraceID:     {1},
		data.AnglerBetaTraceID:      nil,
		data.BullheadAlphaTraceID:   {0, 1},
		data.BullheadBetaTraceID:    {0, 2},
		data.CrosshatchAlphaTraceID: {1},
		data.CrosshatchBetaTraceID:  nil,
	}
	// We only check that the traces have their associated comments. We rely on the other tests
	// to make sure the other fields are correct.
	traceCount := 0
	for _, r := range resp.Digests {
		for _, tr := range r.Traces.Traces {
			traceCount++
			assert.Equal(t, expectedComments[tr.ID], tr.CommentIndices, "trace id %q under digest", tr.ID, r.Digest)
		}
	}
	assert.Equal(t, 6, traceCount, "Not all traces were in the final result")
}

// TestSearchThreeDevicesChangeListSunnyDay covers the case
// where two tryjobs have been run on a given CL and PS, one on the
// angler bot and one on the bullhead bot. The master branch
// looks like in the ThreeDevices data set. The outputs produced are
// Test  |  Device  | Digest
// ----------------------
// Alpha | Angler   | data.AlphaGood1Digest
// Alpha | Bullhead | data.AlphaUntriaged1Digest
// Beta  | Angler   | data.BetaGood1Digest
// Beta  | Bullhead | BetaBrandNewDigest
//
// The user has triaged the data.AlphaUntriaged1Digest as positive
// but BetaBrandNewDigest remains untriaged.
// With this setup, we do a default query (don't show master,
// only untriaged digests) and expect to see only an entry about
// BetaBrandNewDigest.
func TestSearchThreeDevicesChangeListSunnyDay(t *testing.T) {
	unittest.SmallTest(t)

	clID := "1234"
	crs := "gerrit"
	AlphaNowGoodDigest := data.AlphaUntriaged1Digest
	BetaBrandNewDigest := types.Digest("be7a03256511bec3a7453c3186bb2e07")

	mcls := &mock_clstore.Store{}
	mtjs := &mock_tjstore.Store{}
	defer mcls.AssertExpectations(t)
	defer mtjs.AssertExpectations(t)

	mes := makeThreeDevicesExpectationStore()
	var ie expectations.Expectations
	ie.Set(data.AlphaTest, AlphaNowGoodDigest, expectations.Positive)
	addChangeListExpectations(mes, crs, clID, &ie)

	mcls.On("GetPatchSets", testutils.AnyContext, clID).Return([]code_review.PatchSet{
		{
			SystemID:     "first_one",
			ChangeListID: clID,
			Order:        1,
			// All the rest are ignored
		},
		{
			SystemID:     "fourth_one",
			ChangeListID: clID,
			Order:        4,
			// All the rest are ignored
		},
	}, nil).Once() // this should be cached after fetch, as it could be expensive to retrieve.
	mcls.On("System").Return(crs)

	expectedID := tjstore.CombinedPSID{
		CL:  clID,
		CRS: crs,
		PS:  "fourth_one", // we didn't specify a PS, so it goes with the most recent
	}
	anglerGroup := map[string]string{
		"device": data.AnglerDevice,
	}
	bullheadGroup := map[string]string{
		"device": data.BullheadDevice,
	}
	options := map[string]string{
		"ext": "png",
	}

	mtjs.On("GetResults", testutils.AnyContext, expectedID).Return([]tjstore.TryJobResult{
		{
			GroupParams: anglerGroup,
			Options:     options,
			Digest:      data.AlphaGood1Digest,
			ResultParams: map[string]string{
				types.PrimaryKeyField: string(data.AlphaTest),
				types.CorpusField:     "gm",
			},
		},
		{
			GroupParams: bullheadGroup,
			Options:     options,
			Digest:      AlphaNowGoodDigest,
			ResultParams: map[string]string{
				types.PrimaryKeyField: string(data.AlphaTest),
				types.CorpusField:     "gm",
			},
		},
		{
			GroupParams: anglerGroup,
			Options:     options,
			Digest:      data.BetaGood1Digest,
			ResultParams: map[string]string{
				types.PrimaryKeyField: string(data.BetaTest),
				types.CorpusField:     "gm",
			},
		},
		{
			GroupParams: bullheadGroup,
			Options:     options,
			Digest:      BetaBrandNewDigest,
			ResultParams: map[string]string{
				types.PrimaryKeyField: string(data.BetaTest),
				types.CorpusField:     "gm",
			},
		},
	}, nil).Once() // this should be cached after fetch, as it could be expensive to retrieve.

	mds := makeDiffStoreWithNoFailures()
	addDiffData(mds, BetaBrandNewDigest, data.BetaGood1Digest, makeSmallDiffMetric())

	s := New(mds, mes, makeThreeDevicesIndexer(), mcls, mtjs, nil, everythingPublic)

	q := &query.Search{
		ChangeListID:  clID,
		NewCLStore:    true,
		IncludeMaster: false,

		Unt:  true,
		Head: true,

		Metric:   diff.CombinedMetric,
		FRGBAMin: 0,
		FRGBAMax: 255,
		FDiffMax: -1,
		Sort:     query.SortAscending,
	}

	resp, err := s.Search(context.Background(), q)
	require.NoError(t, err)
	require.NotNil(t, resp)
	// make sure the group maps were not mutated.
	assert.Len(t, anglerGroup, 1)
	assert.Len(t, bullheadGroup, 1)
	assert.Len(t, options, 1)

	assert.Equal(t, &frontend.SearchResponse{
		Commits: data.MakeTestCommits(),
		Offset:  0,
		Size:    1,
		Digests: []*frontend.SRDigest{
			{
				Test:   data.BetaTest,
				Digest: BetaBrandNewDigest,
				Status: "untriaged",
				ParamSet: map[string][]string{
					"device":              {data.BullheadDevice},
					types.PrimaryKeyField: {string(data.BetaTest)},
					types.CorpusField:     {"gm"},
					"ext":                 {"png"},
				},
				Traces: &frontend.TraceGroup{
					TileSize: 3,
					Traces:   []frontend.Trace{},
					Digests: []frontend.DigestStatus{
						{
							Digest: BetaBrandNewDigest,
							Status: "untriaged",
						},
					},
				},
				ClosestRef: common.PositiveRef,
				RefDiffs: map[common.RefClosest]*frontend.SRDiffDigest{
					common.PositiveRef: {
						DiffMetrics: makeSmallDiffMetric(),
						Digest:      data.BetaGood1Digest,
						Status:      "positive",
						ParamSet: map[string][]string{
							"device":              {data.AnglerDevice, data.BullheadDevice},
							types.PrimaryKeyField: {string(data.BetaTest)},
							types.CorpusField:     {"gm"},
							// Note: the data from three_devices lacks an "ext" entry, so
							// we don't see one here
						},
						OccurrencesInTile: 6,
					},
					common.NegativeRef: nil,
				},
			},
		},
	}, resp)

	// Validate that we cache the .*Store values in two quick responses.
	_, err = s.Search(context.Background(), q)
	require.NoError(t, err)
}

func TestDigestDetailsThreeDevicesSunnyDay(t *testing.T) {
	unittest.SmallTest(t)

	const digestWeWantDetailsAbout = data.AlphaGood1Digest
	const testWeWantDetailsAbout = data.AlphaTest

	mds := makeDiffStoreWithNoFailures()
	// Note: If a digest is compared to itself, it is removed from the return value, so we use nil.
	addDiffData(mds, digestWeWantDetailsAbout, data.AlphaGood1Digest, nil)
	addDiffData(mds, digestWeWantDetailsAbout, data.AlphaBad1Digest, makeBigDiffMetric())

	s := New(mds, makeThreeDevicesExpectationStore(), makeThreeDevicesIndexer(), nil, nil, emptyCommentStore(), everythingPublic)

	details, err := s.GetDigestDetails(context.Background(), testWeWantDetailsAbout, digestWeWantDetailsAbout, "", "")
	require.NoError(t, err)
	assert.Equal(t, &frontend.DigestDetails{
		Commits: data.MakeTestCommits(),
		Digest: &frontend.SRDigest{
			Test:   testWeWantDetailsAbout,
			Digest: digestWeWantDetailsAbout,
			Status: "positive",
			ParamSet: map[string][]string{
				"device":              {data.AnglerDevice, data.CrosshatchDevice},
				types.PrimaryKeyField: {string(data.AlphaTest)},
				types.CorpusField:     {"gm"},
			},
			Traces: &frontend.TraceGroup{
				TileSize: 3, // 3 commits in tile
				Traces: []frontend.Trace{ // the digest we care about appears in two traces
					{
						Data: []int{1, 1, 0},
						ID:   data.AnglerAlphaTraceID,
						Params: map[string]string{
							"device":              data.AnglerDevice,
							types.PrimaryKeyField: string(data.AlphaTest),
							types.CorpusField:     "gm",
						},
					},
					{
						Data: []int{1, 1, 0},
						ID:   data.CrosshatchAlphaTraceID,
						Params: map[string]string{
							"device":              data.CrosshatchDevice,
							types.PrimaryKeyField: string(data.AlphaTest),
							types.CorpusField:     "gm",
						},
					},
				},
				Digests: []frontend.DigestStatus{
					{
						Digest: data.AlphaGood1Digest,
						Status: "positive",
					},
					{
						Digest: data.AlphaBad1Digest,
						Status: "negative",
					},
				},
			},
			ClosestRef: common.NegativeRef,
			RefDiffs: map[common.RefClosest]*frontend.SRDiffDigest{
				common.PositiveRef: nil,
				common.NegativeRef: {
					DiffMetrics: makeBigDiffMetric(),
					Digest:      data.AlphaBad1Digest,
					Status:      "negative",
					ParamSet: map[string][]string{
						"device":              {data.AnglerDevice, data.BullheadDevice, data.CrosshatchDevice},
						types.PrimaryKeyField: {string(data.AlphaTest)},
						types.CorpusField:     {"gm"},
					},
					OccurrencesInTile: 6,
				},
			},
		},
	}, details)
}

func TestDigestDetailsThreeDevicesChangeList(t *testing.T) {
	unittest.SmallTest(t)

	const digestWeWantDetailsAbout = data.AlphaGood1Digest
	const testWeWantDetailsAbout = data.AlphaTest
	const testCLID = "abc12345"
	const testCRS = "gerritHub"

	mes := makeThreeDevicesExpectationStore()
	// Mock out some ChangeList expectations in which the digest we care about is negative
	var ie expectations.Expectations
	ie.Set(testWeWantDetailsAbout, digestWeWantDetailsAbout, expectations.Negative)
	addChangeListExpectations(mes, testCRS, testCLID, &ie)

	mds := makeDiffStoreWithNoFailures()
	// There are no positive digests with which to compare
	// Negative match. Note If a digest is compared to itself, it is removed from the return value.
	mds.On("Get", testutils.AnyContext, digestWeWantDetailsAbout, types.DigestSlice{digestWeWantDetailsAbout, data.AlphaBad1Digest}).
		Return(map[types.Digest]*diff.DiffMetrics{
			data.AlphaBad1Digest: makeBigDiffMetric(),
		}, nil)

	s := New(mds, mes, makeThreeDevicesIndexer(), nil, nil, emptyCommentStore(), everythingPublic)

	details, err := s.GetDigestDetails(context.Background(), testWeWantDetailsAbout, digestWeWantDetailsAbout, testCLID, testCRS)
	require.NoError(t, err)
	assert.Equal(t, details.Digest.Status, expectations.Negative.String())
}

// TestDigestDetailsThreeDevicesOldDigest represents the scenario in which a user is requesting
// data about a digest that just went off the tile.
func TestDigestDetailsThreeDevicesOldDigest(t *testing.T) {
	unittest.SmallTest(t)

	const digestWeWantDetailsAbout = types.Digest("digest-too-old")
	const testWeWantDetailsAbout = data.BetaTest

	mds := makeDiffStoreWithNoFailures()
	addDiffData(mds, digestWeWantDetailsAbout, data.BetaGood1Digest, makeSmallDiffMetric())

	s := New(mds, makeThreeDevicesExpectationStore(), makeThreeDevicesIndexer(), nil, nil, nil, everythingPublic)

	d, err := s.GetDigestDetails(context.Background(), testWeWantDetailsAbout, digestWeWantDetailsAbout, "", "")
	require.NoError(t, err)
	// spot check is fine for this test because other tests do a more thorough check of the
	// whole struct.
	assert.Equal(t, digestWeWantDetailsAbout, d.Digest.Digest)
	assert.Equal(t, testWeWantDetailsAbout, d.Digest.Test)
	assert.Equal(t, map[common.RefClosest]*frontend.SRDiffDigest{
		common.PositiveRef: {
			DiffMetrics: makeSmallDiffMetric(),
			Digest:      data.BetaGood1Digest,
			Status:      "positive",
			ParamSet: paramtools.ParamSet{
				"device":              []string{data.AnglerDevice, data.BullheadDevice},
				types.PrimaryKeyField: []string{string(data.BetaTest)},
				types.CorpusField:     []string{"gm"},
			},
			OccurrencesInTile: 6,
		},
		common.NegativeRef: nil,
	}, d.Digest.RefDiffs)
}

// TestDigestDetailsThreeDevicesOldDigest represents the scenario in which a user is requesting
// data about a digest that never existed. In the past, when this has happened, it has broken
// Gold until that digest went away (e.g. because a bot only uploaded a subset of images).
// Therefore, we shouldn't error the search request, because it could break all searches for
// untriaged digests.
func TestDigestDetailsThreeDevicesBadDigest(t *testing.T) {
	unittest.SmallTest(t)

	const digestWeWantDetailsAbout = types.Digest("unknown-digest")
	const testWeWantDetailsAbout = data.BetaTest

	mds := makeDiffStoreWithNoFailures()
	mds.On("Get", testutils.AnyContext, digestWeWantDetailsAbout, types.DigestSlice{data.BetaGood1Digest}).Return(nil, errors.New("invalid digest"))

	s := New(mds, makeThreeDevicesExpectationStore(), makeThreeDevicesIndexer(), nil, nil, nil, everythingPublic)

	r, err := s.GetDigestDetails(context.Background(), testWeWantDetailsAbout, digestWeWantDetailsAbout, "", "")
	require.NoError(t, err)
	// Since we couldn't find the digest, we have nothing to compare against.
	assert.Equal(t, r.Digest.Digest, digestWeWantDetailsAbout)
	assert.Equal(t, r.Digest.ClosestRef, common.NoRef)
}

func TestDigestDetailsThreeDevicesBadTest(t *testing.T) {
	unittest.SmallTest(t)

	const digestWeWantDetailsAbout = data.AlphaGood1Digest
	const testWeWantDetailsAbout = types.TestName("invalid test")

	s := New(nil, nil, makeThreeDevicesIndexer(), nil, nil, nil, everythingPublic)

	_, err := s.GetDigestDetails(context.Background(), testWeWantDetailsAbout, digestWeWantDetailsAbout, "", "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unknown")
}

func TestDigestDetailsThreeDevicesBadTestAndDigest(t *testing.T) {
	unittest.SmallTest(t)

	const digestWeWantDetailsAbout = types.Digest("invalid digest")
	const testWeWantDetailsAbout = types.TestName("invalid test")

	s := New(nil, nil, makeThreeDevicesIndexer(), nil, nil, nil, everythingPublic)

	_, err := s.GetDigestDetails(context.Background(), testWeWantDetailsAbout, digestWeWantDetailsAbout, "", "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unknown")
}

func TestDiffDigestsSunnyDay(t *testing.T) {
	unittest.SmallTest(t)

	const testWeWantDetailsAbout = data.AlphaTest
	const leftDigest = data.AlphaUntriaged1Digest
	const rightDigest = data.AlphaGood1Digest

	mds := makeDiffStoreWithNoFailures()
	addDiffData(mds, leftDigest, rightDigest, makeSmallDiffMetric())

	s := New(mds, makeThreeDevicesExpectationStore(), makeThreeDevicesIndexer(), nil, nil, nil, everythingPublic)

	cd, err := s.DiffDigests(context.Background(), testWeWantDetailsAbout, leftDigest, rightDigest, "", "")
	require.NoError(t, err)
	assert.Equal(t, &frontend.DigestComparison{
		Left: &frontend.SRDigest{
			Test:   testWeWantDetailsAbout,
			Digest: leftDigest,
			Status: expectations.Untriaged.String(),
			ParamSet: paramtools.ParamSet{
				"device":              []string{data.BullheadDevice},
				types.PrimaryKeyField: []string{string(data.AlphaTest)},
				types.CorpusField:     []string{"gm"},
			},
		},
		Right: &frontend.SRDiffDigest{
			Digest:      rightDigest,
			Status:      expectations.Positive.String(),
			DiffMetrics: makeSmallDiffMetric(),
			ParamSet: paramtools.ParamSet{
				"device":              []string{data.AnglerDevice, data.CrosshatchDevice},
				types.PrimaryKeyField: []string{string(data.AlphaTest)},
				types.CorpusField:     []string{"gm"},
			},
		},
	}, cd)
}

func TestDiffDigestsChangeList(t *testing.T) {
	unittest.SmallTest(t)

	const testWeWantDetailsAbout = data.AlphaTest
	const leftDigest = data.AlphaUntriaged1Digest
	const rightDigest = data.AlphaGood1Digest
	const clID = "abc12354"
	const crs = "gerritHub"

	mes := makeThreeDevicesExpectationStore()
	var ie expectations.Expectations
	ie.Set(data.AlphaTest, leftDigest, expectations.Negative)
	addChangeListExpectations(mes, crs, clID, &ie)

	mds := makeDiffStoreWithNoFailures()
	addDiffData(mds, leftDigest, rightDigest, makeSmallDiffMetric())

	s := New(mds, mes, makeThreeDevicesIndexer(), nil, nil, nil, everythingPublic)

	cd, err := s.DiffDigests(context.Background(), testWeWantDetailsAbout, leftDigest, rightDigest, clID, crs)
	require.NoError(t, err)
	assert.Equal(t, cd.Left.Status, expectations.Negative.String())
}

// TestUntriagedUnignoredTryJobExclusiveDigestsSunnyDay models the case where a set of TryJobs has
// produced five digests that were "untriaged on master" (and one good digest). We are testing that
// we can properly deduce which are untriaged, "newly seen" and unignored. One of these untriaged
// digests was already seen on master (data.AlphaUntriaged1Digest), one was already triaged negative
// for this CL (gammaNegativeTryJobDigest), and one trace matched an ignore rule (deltaIgnoredTryJobDigest). Thus,
// We only expect tjUntriagedAlpha and tjUntriagedBeta to be reported to us.
func TestUntriagedUnignoredTryJobExclusiveDigestsSunnyDay(t *testing.T) {
	unittest.SmallTest(t)

	const clID = "44474"
	const crs = "github"
	expectedID := tjstore.CombinedPSID{
		CL:  clID,
		CRS: crs,
		PS:  "abcdef",
	}

	const alphaUntriagedTryJobDigest = types.Digest("aaaa65e567de97c8a62918401731c7ec")
	const betaUntriagedTryJobDigest = types.Digest("bbbb34f7c915a1ac3a5ba524c741946c")
	const gammaNegativeTryJobDigest = types.Digest("cccc41bf4584e51be99e423707157277")
	const deltaIgnoredTryJobDigest = types.Digest("dddd84e51be99e42370715727765e563")

	mi := &mock_index.IndexSource{}
	mtjs := &mock_tjstore.Store{}

	// Set up the expectations such that for this CL, we have one extra expectation - marking
	// gammaNegativeTryJobDigest negative (it would be untriaged on master).
	mes := makeThreeDevicesExpectationStore()
	var ie expectations.Expectations
	ie.Set(data.AlphaTest, gammaNegativeTryJobDigest, expectations.Negative)
	addChangeListExpectations(mes, crs, clID, &ie)

	cpxTile := types.NewComplexTile(data.MakeTestTile())
	reduced := data.MakeTestTile()
	delete(reduced.Traces, data.BullheadBetaTraceID)
	// The following rule exclusively matches BullheadBetaTraceID, for which the tryjob produced
	// deltaIgnoredTryJobDigest
	cpxTile.SetIgnoreRules(reduced, paramtools.ParamMatcher{
		{
			"device":              []string{data.BullheadDevice},
			types.PrimaryKeyField: []string{string(data.BetaTest)},
		},
	})
	dc := digest_counter.New(data.MakeTestTile())
	fis, err := indexer.SearchIndexForTesting(cpxTile, [2]digest_counter.DigestCounter{dc, dc}, [2]paramsets.ParamSummary{}, mes, nil)
	require.NoError(t, err)
	mi.On("GetIndex").Return(fis)

	anglerGroup := map[string]string{
		"device": data.AnglerDevice,
	}
	bullheadGroup := map[string]string{
		"device": data.BullheadDevice,
	}
	crosshatchGroup := map[string]string{
		"device": data.CrosshatchDevice,
	}
	options := map[string]string{
		"ext": "png",
	}
	mtjs.On("GetResults", testutils.AnyContext, expectedID).Return([]tjstore.TryJobResult{
		{
			GroupParams: anglerGroup,
			Options:     options,
			Digest:      betaUntriagedTryJobDigest,
			ResultParams: map[string]string{
				types.PrimaryKeyField: string(data.AlphaTest),
				types.CorpusField:     "gm",
			},
		},
		{
			GroupParams: bullheadGroup,
			Options:     options,
			Digest:      data.AlphaUntriaged1Digest,
			ResultParams: map[string]string{
				types.PrimaryKeyField: string(data.AlphaTest),
				types.CorpusField:     "gm",
			},
		},
		{
			GroupParams: anglerGroup,
			Options:     options,
			Digest:      alphaUntriagedTryJobDigest,
			ResultParams: map[string]string{
				types.PrimaryKeyField: string(data.BetaTest),
				types.CorpusField:     "gm",
			},
		},
		{
			GroupParams: bullheadGroup,
			Options:     options,
			Digest:      deltaIgnoredTryJobDigest,
			ResultParams: map[string]string{
				types.PrimaryKeyField: string(data.BetaTest),
				types.CorpusField:     "gm",
			},
		},
		{
			GroupParams: crosshatchGroup,
			Options:     options,
			Digest:      gammaNegativeTryJobDigest,
			ResultParams: map[string]string{
				types.PrimaryKeyField: string(data.AlphaTest),
				types.CorpusField:     "gm",
			},
		},
		{
			GroupParams: crosshatchGroup,
			Options:     options,
			Digest:      data.BetaGood1Digest,
			ResultParams: map[string]string{
				types.PrimaryKeyField: string(data.BetaTest),
				types.CorpusField:     "gm",
			},
		},
	}, nil).Once()

	s := New(nil, mes, mi, nil, mtjs, nil, everythingPublic)

	dl, err := s.UntriagedUnignoredTryJobExclusiveDigests(context.Background(), expectedID)
	require.NoError(t, err)
	assert.Equal(t, []types.Digest{alphaUntriagedTryJobDigest, betaUntriagedTryJobDigest}, dl.Digests)
}

// TestGetDrawableTraces_DigestIndicesAreCorrect tests that we generate the output required to draw
// the trace graphs correctly, especially when dealing with many digests or missing digests.
func TestGetDrawableTraces_DigestIndicesAreCorrect(t *testing.T) {
	unittest.SmallTest(t)
	// Add some shorthand aliases for easier-to-read test inputs.
	const mm = types.MissingDigest
	const mdi = missingDigestIndex
	// These constants are not actual md5 digests, but that's ok for the purposes of this test -
	// any string constants will do.
	const d0, d1, d2, d3, d4 = types.Digest("a"), types.Digest("b"), types.Digest("c"), types.Digest("d"), types.Digest("e")

	test := func(desc string, inputDigests []types.Digest, expectedData []int) {
		// stubClassifier returns Positive for everything. For the purposes of drawing traces,
		// don't actually care about the expectations.
		stubClassifier := &mock_expectations.Classifier{}
		stubClassifier.On("Classification", mock.Anything, mock.Anything).Return(expectations.Positive)
		t.Run(desc, func(t *testing.T) {
			s := SearchImpl{}
			traces := map[tiling.TraceID]*types.GoldenTrace{
				"not-a-real-trace-id-and-that's-ok": {
					Digests: inputDigests,
					// Keys can be omitted because they are not read here,
				},
			}
			rv := s.getDrawableTraces("whatever", d0, len(inputDigests)-1, stubClassifier, traces, nil)
			require.Len(t, rv.Traces, 1)
			assert.Equal(t, expectedData, rv.Traces[0].Data)
		})
	}

	test("several distinct digests",
		[]types.Digest{d4, d3, d2, d1, d0},
		[]int{4, 3, 2, 1, 0})
	// index 1 represents the first digest, starting at head, that doesn't match the "digest of
	// focus", which for these tests is d0. For convenience, in all the other sub-tests, the index
	// on the constants matches the expected index.
	test("several distinct digests, ordered by proximity to head",
		[]types.Digest{d1, d2, d3, d4, d0},
		[]int{4, 3, 2, 1, 0})
	test("missing digests",
		[]types.Digest{mm, d1, mm, d0, mm},
		[]int{mdi, 1, mdi, 0, mdi})
	test("multiple missing digest in a row",
		[]types.Digest{mm, mm, mm, d1, d1, mm, mm, mm, d0, mm, mm},
		[]int{mdi, mdi, mdi, 1, 1, mdi, mdi, mdi, 0, mdi, mdi})
	test("all the same",
		[]types.Digest{d0, d0, d0, d0, d0, d0, d0},
		[]int{0, 0, 0, 0, 0, 0, 0})
	test("d0 not at head",
		[]types.Digest{d0, d0, d0, d1, d2, d1},
		[]int{0, 0, 0, 1, 2, 1})
	// At a certain point, we lump distinct digests together. Currently this is after we have seen
	// 9 distinct digests (starting at head).
	test("too many distinct digests",
		[]types.Digest{"dA", "d9", "d8", "d7", "d6", "d5", d4, d3, d2, d1, d0},
		[]int{8, 8, 8, 7, 6, 5, 4, 3, 2, 1, 0})
}

var everythingPublic = paramtools.ParamSet{}

// makeThreeDevicesIndexer returns an IndexSource that returns the result of makeThreeDevicesIndex.
func makeThreeDevicesIndexer() indexer.IndexSource {
	mi := &mock_index.IndexSource{}

	fis := makeThreeDevicesIndex()
	mi.On("GetIndex").Return(fis)
	return mi
}

// makeThreeDevicesIndex returns a search index corresponding to the three_devices_data
// (which currently has nothing ignored).
func makeThreeDevicesIndex() *indexer.SearchIndex {
	cpxTile := types.NewComplexTile(data.MakeTestTile())
	dc := digest_counter.New(data.MakeTestTile())
	ps := paramsets.NewParamSummary(data.MakeTestTile(), dc)
	si, err := indexer.SearchIndexForTesting(cpxTile, [2]digest_counter.DigestCounter{dc, dc}, [2]paramsets.ParamSummary{ps, ps}, nil, nil)
	if err != nil {
		// Something is horribly broken with our test data/setup
		panic(err.Error())
	}
	return si
}

func makeThreeDevicesExpectationStore() *mock_expectations.Store {
	mes := &mock_expectations.Store{}
	mes.On("Get", testutils.AnyContext).Return(data.MakeTestExpectations(), nil)
	return mes
}

func addChangeListExpectations(mes *mock_expectations.Store, crs string, clID string, issueExp *expectations.Expectations) {
	issueStore := &mock_expectations.Store{}
	mes.On("ForChangeList", clID, crs).Return(issueStore, nil)
	issueStore.On("Get", testutils.AnyContext).Return(issueExp, nil)
}

func makeDiffStoreWithNoFailures() *mock_diffstore.DiffStore {
	mds := &mock_diffstore.DiffStore{}
	mds.On("UnavailableDigests", testutils.AnyContext).Return(map[types.Digest]*diff.DigestFailure{}, nil)
	return mds
}

func addDiffData(mds *mock_diffstore.DiffStore, left types.Digest, right types.Digest, metric *diff.DiffMetrics) {
	if metric == nil {
		// empty map is expected instead of a nil entry
		mds.On("Get", testutils.AnyContext, left, types.DigestSlice{right}).
			Return(map[types.Digest]*diff.DiffMetrics{}, nil)
	} else {
		mds.On("Get", testutils.AnyContext, left, types.DigestSlice{right}).
			Return(map[types.Digest]*diff.DiffMetrics{
				right: metric,
			}, nil)
	}
}

// This is arbitrary data.
func makeSmallDiffMetric() *diff.DiffMetrics {
	return &diff.DiffMetrics{
		NumDiffPixels:    8,
		PixelDiffPercent: 0.02,
		MaxRGBADiffs:     [4]int{0, 48, 12, 0},
		DimDiffer:        false,
		Diffs: map[string]float32{
			diff.CombinedMetric: 0.0005,
			"percent":           0.02,
			"pixel":             8,
		},
	}
}

func makeBigDiffMetric() *diff.DiffMetrics {
	return &diff.DiffMetrics{
		NumDiffPixels:    88812,
		PixelDiffPercent: 98.68,
		MaxRGBADiffs:     [4]int{102, 51, 13, 0},
		DimDiffer:        true,
		Diffs: map[string]float32{
			diff.CombinedMetric: 4.7,
			"percent":           98.68,
			"pixel":             88812,
		},
	}
}

func emptyCommentStore() comment.Store {
	mcs := &mock_comment.Store{}
	mcs.On("ListComments", testutils.AnyContext).Return(nil, nil)
	return mcs
}

// makeStubDiffStore returns a diffstore that returns the small diff metric for every call to Get.
func makeStubDiffStore() *mock_diffstore.DiffStore {
	mds := &mock_diffstore.DiffStore{}
	mds.On("UnavailableDigests", testutils.AnyContext).Return(map[types.Digest]*diff.DigestFailure{}, nil)
	mds.On("Get", testutils.AnyContext, mock.Anything, mock.Anything).Return(func(_ context.Context, _ types.Digest, rights types.DigestSlice) map[types.Digest]*diff.DiffMetrics {
		rv := make(map[types.Digest]*diff.DiffMetrics, len(rights))
		for _, right := range rights {
			rv[right] = makeSmallDiffMetric()
		}
		return rv
	}, nil)
	return mds
}
