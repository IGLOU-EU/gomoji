#!/bin/bash

# Internal traitement
declare -a emojis
declare -A cat key

# Global output
declare HASH_LT EMOJIS EMOJIS_NB CAT CAT_NB KEY KEY_NB

# Utils var
dirScript="$(dirname "$(readlink -fn "$0")")"
dirRoot="$(dirname "$dirScript")"

fileTpl="$dirScript/data.tpl.go"
fileOut="$dirRoot/data.go"
fileLatest="$dirScript/latest.source"

urlEmoji="https://unicode.org/emoji/charts/emoji-list.html"

# Source previous build
# shellcheck source=latest.source
source "$fileLatest"

error() {
	printf \
		'%s	| %s:%s	| %s: %b\n' \
		"$(date -Iminutes)" "$(basename "$0")" "${BASH_LINENO[0]}" \
		"${2:-ERROR}" "$1" >&2
}

fatal() {
	error "$1" "FATAL"
	exit 1
}

url_get() {
	local _content
	local _retry=5

	for ((i = 0; i < _retry; i++)); do
		if _content="$(curl -f -sS "$1")"; then
			break
		fi

		error "can't get '$1' _retry $((i + 1))"
		sleep 1
	done

	echo "$_content"
}

main() {
	local _content _hash
	_content="$(url_get "$urlEmoji" | tr -d '\r\n')"
	_hash="$(printf '%s' "$_content" | shasum -a 256 | awk '{print $1}')"

	# Check if there is something new
	if [[ $_hash == "$HASH_LT" ]]; then
		exit 0
	fi

	# Syntax analysis and data processing
	local catActive
	while read -r line; do
		case "$line" in

		*\"bighead\"*)
			catActive="$(htmlq -t 'a' <<<"$line")"
			cat[$catActive]=0
			;;

		*\"imga\"*)
			local kw
			local _buff keywords

			keywords=''
			kw="$(htmlq -t 'div:nth-child(5)' <<<"$line")"
			while read -r k; do
				if [[ -z $keywords ]]; then
					keywords="\"$k\""
				else
					keywords="$keywords,\"$k\""
				fi

				k="$(sed 's/*/|*/g' <<<"$k")"
				key[$k]=$((key[$k] + 1))
			done <<<"$(sed 's/ | /\n/g' <<<"$kw")"

			cat[$catActive]=$((cat[$catActive] + 1))

			printf -v _buff '{"%s", "%s", "%s", []string{%s}},' \
				"$(htmlq -a alt 'img.imga' <<<"$line")" \
				"$(htmlq -t 'div:nth-child(4)' <<<"$line")" \
				"$catActive" \
				"$keywords"

			emojis+=("$_buff")
			;;
		esac
	done <<<"$(htmlq '.main > table tr' <<<"$_content" | sed 's/⊛ //g; s/td/div/g')"

	# Data formating
	EMOJIS="$(sed 's/&/\\&/g' <<<"${emojis[*]}")"
	EMOJIS_NB=${#emojis[@]}

	CAT_NB=${#cat[@]}
	for c in "${!cat[@]}"; do
		CAT="$(printf '{"%s", %d},' "$c" "${cat[$c]}") $CAT"
	done

	KEY_NB=${#key[@]}
	for k in "${!key[@]}"; do
		KEY="$KEY $(printf '{"%s", %d},' "$(sed 's/|\*/*/g' <<<"$k")" "${key[$k]}")"
	done

	# Write to the output
	export EMOJIS_NB CAT CAT_NB KEY KEY_NB
	envsubst <"$fileTpl" >"$fileOut"

	sed -i -f <(printf '%s\n' "s^%EMOJIS^$EMOJIS^g") "$fileOut"

	# Wrtie the latest file
	printf	'#!/bin/bash\n# shellcheck disable=SC2034\nDATE_LT="%s"\nHASH_LT="%s"\nEMOJIS_NB=%d\nCAT_NB=%d\nKEY_NB=%d' \
		"$(date -u --iso-8601)" "$_hash" "$EMOJIS_NB" "$CAT_NB" "$KEY_NB" \
		> "$fileLatest"
}

main
