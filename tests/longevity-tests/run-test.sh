#!/bin/bash
DEFAULT_TESTCASES=("tests/manual-test-cases/Group14-Longevity" "tests/manual-test-cases/Group11-Stress")
testcases=("${@:-${DEFAULT_TESTCASES[@]}}")
pybot "${testcases[@]}"
