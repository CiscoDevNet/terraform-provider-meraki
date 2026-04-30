#!/usr/bin/env bash
set -euo pipefail

run_all() {
  echo "run_tests=true" >> "$GITHUB_OUTPUT"
  echo "run_all=true" >> "$GITHUB_OUTPUT"
  echo "test_pattern=" >> "$GITHUB_OUTPUT"
  echo "=> Running ALL tests"
  exit 0
}

skip_tests() {
  echo "run_tests=false" >> "$GITHUB_OUTPUT"
  echo "run_all=false" >> "$GITHUB_OUTPUT"
  echo "test_pattern=" >> "$GITHUB_OUTPUT"
  echo "=> No tests to run"
  exit 0
}

if [[ "${FULL_TEST:-false}" == "true" ]]; then
  echo "Full test requested via workflow_dispatch"
  run_all
fi

if [[ "${GITHUB_REF:-}" == refs/tags/* ]]; then
  echo "Tag push detected: ${GITHUB_REF} — running full suite"
  run_all
fi

if [[ "$GITHUB_EVENT_NAME" == "pull_request" ]]; then
  BASE_SHA="${GITHUB_BASE_REF}"
  git fetch --depth=1 origin "$BASE_SHA" 2>/dev/null || true
  DIFF_BASE="origin/${BASE_SHA}"
else
  DIFF_BASE="HEAD~1"
fi

CHANGED_FILES=$(git diff --name-only "$DIFF_BASE" HEAD 2>/dev/null || git diff --name-only HEAD~1 HEAD 2>/dev/null || echo "")

if [[ -z "$CHANGED_FILES" ]]; then
  echo "No changed files detected"
  skip_tests
fi

echo "Changed files:"
echo "$CHANGED_FILES"
echo ""

RUN_ALL_PATTERNS=(
  "^go\.mod$"
  "^go\.sum$"
  "^main\.go$"
  "^internal/provider/provider\.go$"
  "^internal/provider/provider_test\.go$"
  "^internal/provider/helpers"
)

for pattern in "${RUN_ALL_PATTERNS[@]}"; do
  if echo "$CHANGED_FILES" | grep -qE "$pattern"; then
    echo "Matched run-all pattern: $pattern"
    run_all
  fi
done

PROVIDER_FILES=$(echo "$CHANGED_FILES" | grep -E '^internal/provider/.*\.go$' || true)

if [[ -z "$PROVIDER_FILES" ]]; then
  echo "No provider Go files changed"
  skip_tests
fi

declare -A TEST_FILES_MAP

for file in $PROVIDER_FILES; do
  basename=$(basename "$file")

  case "$basename" in
    resource_meraki_*_test.go)
      TEST_FILES_MAP["internal/provider/$basename"]=1
      ;;
    data_source_meraki_*_test.go)
      TEST_FILES_MAP["internal/provider/$basename"]=1
      ;;
    resource_meraki_*.go)
      test_file="${basename%.go}_test.go"
      TEST_FILES_MAP["internal/provider/$test_file"]=1
      ;;
    data_source_meraki_*.go)
      test_file="${basename%.go}_test.go"
      TEST_FILES_MAP["internal/provider/$test_file"]=1
      ;;
    model_resource_meraki_*.go)
      resource_name="${basename#model_}"
      test_file="${resource_name%.go}_test.go"
      TEST_FILES_MAP["internal/provider/$test_file"]=1
      ;;
    model_data_source_meraki_*.go)
      ds_name="${basename#model_}"
      test_file="${ds_name%.go}_test.go"
      TEST_FILES_MAP["internal/provider/$test_file"]=1
      ;;
    model_meraki_*.go)
      resource_name="${basename#model_}"
      res_test="resource_${resource_name%.go}_test.go"
      ds_test="data_source_${resource_name%.go}_test.go"
      TEST_FILES_MAP["internal/provider/$res_test"]=1
      TEST_FILES_MAP["internal/provider/$ds_test"]=1
      ;;
    *)
      echo "Unrecognized provider file pattern: $basename — running all tests"
      run_all
      ;;
  esac
done

TEST_FUNCS=()
for test_file in "${!TEST_FILES_MAP[@]}"; do
  if [[ -f "$test_file" ]]; then
    while IFS= read -r func_name; do
      TEST_FUNCS+=("$func_name")
    done < <(grep -oE '^func (TestAcc[A-Za-z0-9_]+)' "$test_file" | awk '{print $2}')
  else
    echo "Warning: expected test file not found: $test_file"
  fi
done

if [[ ${#TEST_FUNCS[@]} -eq 0 ]]; then
  echo "No test functions found in target files"
  skip_tests
fi

PATTERN=$(printf "%s|" "${TEST_FUNCS[@]}")
PATTERN="^(${PATTERN%|})$"

echo ""
echo "=> Running ${#TEST_FUNCS[@]} test(s)"
echo "=> Pattern: $PATTERN"

echo "run_tests=true" >> "$GITHUB_OUTPUT"
echo "run_all=false" >> "$GITHUB_OUTPUT"
echo "test_pattern=$PATTERN" >> "$GITHUB_OUTPUT"
