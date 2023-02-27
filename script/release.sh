#!/bin/bash

# Utils var
dirScript="$(dirname "$(readlink -fn "$0")")"
dirRoot="$(dirname "$dirScript")"

fileReadme="$dirRoot/README.md"
fileLatest="$dirScript/latest.source"

# Source previous build
# shellcheck source=latest.source
source "$fileLatest"

pt_date="s/^(\*\*Date\*\*: )\`.*\`/\1\`$DATE_LT\`/g"
pt_hash="s/^(\*\*Release ID\*\*: )\`.*\`/\1\`$HASH_LT\`/g"
pt_emoji="s/^(- \*\*Emoji\*\*: )\`.*\`/\1\`$EMOJIS_NB\`/g"
pt_cat="s/^(- \*\*Category\*\*: )\`.*\`/\1\`$CAT_NB\`/g"
pt_key="s/^(- \*\*Keywords\*\*: )\`.*\`/\1\`$KEY_NB\`/g"

content="$(cat "$fileReadme")"
sed -E "$pt_date; $pt_hash; $pt_emoji; $pt_cat; $pt_key" \
    > "$fileReadme" <<< "$content"
