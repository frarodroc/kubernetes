sleep 1
if [ "$FAIL" == true ]
then
  echo $(date) Helth check failed
  exit 1
else
  echo $(date) Helth check passed
  exit 0
fi
