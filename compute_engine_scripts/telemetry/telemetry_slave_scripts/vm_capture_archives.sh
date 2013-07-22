#!/bin/bash
#
# Runs all steps in vm_setup_slave.sh, executes record_wpr and copies the
# created archives to Google Storage.
#
# The script should be run from the skia-telemetry-slave GCE instance's
# /home/default/skia-repo/buildbot/compute_engine_scripts/telemetry/telemetry_slave_scripts
# directory.
#
# Copyright 2013 Google Inc. All Rights Reserved.
# Author: rmistry@google.com (Ravi Mistry)


if [ $# -ne 1 ]; then
  echo
  echo "Usage: `basename $0` 1"
  echo
  echo "The first argument is the slave_num of this telemetry slave."
  echo
  exit 1
fi

SLAVE_NUM=$1

source ../vm_config.sh
source vm_utils.sh

create_worker_file $RECORD_WPR_ACTIVITY

source vm_setup_slave.sh

# Create the webpages_archive directory.
mkdir -p /home/default/storage/webpages_archive/
rm -rf /home/default/storage/webpages_archive/*

for page_set in /home/default/storage/page_sets/*; do
  if [[ -f $page_set ]]; then
    echo "========== Processing $page_set =========="
    check_and_run_xvfb
    DISPLAY=:0 timeout 600 tools/perf/record_wpr --extra-browser-args=--disable-setuid-sandbox --browser=system $page_set
    if [ $? -eq 124 ]; then
      echo "========== $page_set timed out! =========="
    else
      echo "========== Done with $page_set =========="
    fi
  fi
done

# Copy the webpages_archive directory to Google Storage.
gsutil rm -R gs://chromium-skia-gm/telemetry/webpages_archive/slave$SLAVE_NUM/*
gsutil cp /home/default/storage/webpages_archive/* \
  gs://chromium-skia-gm/telemetry/webpages_archive/slave$SLAVE_NUM/

# Create a TIMESTAMP file and copy it to Google Storage.
TIMESTAMP=`date +%s`
echo $TIMESTAMP > /tmp/$TIMESTAMP
cp /tmp/$TIMESTAMP /home/default/storage/webpages_archive/TIMESTAMP
gsutil cp /tmp/$TIMESTAMP gs://chromium-skia-gm/telemetry/webpages_archive/slave$SLAVE_NUM/TIMESTAMP
rm /tmp/$TIMESTAMP

delete_worker_file $RECORD_WPR_ACTIVITY
