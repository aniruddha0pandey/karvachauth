#!/bin/bash

set -a
[ -f .env ] && . .env
set +a

help_text="
NAME:
	karvachauth - A command line utility to add collaborators to GitHub repositories

USAGE:
	karvachauth [global options] command [command options] [arguments...]

VERSION:
	0.0.1

COMMANDS:
	login			Enter username password to OAuth token
	repos			list all current repositories
	collab			list all current collabrators
	add				add collabrators to repositories
	remove			remove collabrators from repositories

GLOBAL OPTIONS:
	--help, -h		Shows a list of commands or help for one command
	--version, -v	Print the version"

case $1 in

	login )
		if [ -z $ACCESS_TOKEN ]
		then
			if [ -z $2 ]
			then
				echo "[0;31mInvalid Access Token[0m"
				echo "$help_text"
			else
				ACCESS_TOKEN=$2
				echo "[1;33mWaiting for authentication...[0m"
				res=$(curl -s -H "Authorization: token $ACCESS_TOKEN" https://api.github.com/user)
				username=$(echo $res | ./jq-win64 -r '.login')
				if [[ $username == "null" ]]
				then
					echo "[0;31mInvalid Client Credentials ;([0m"
				else
					echo "[0;32mSuccess! Logged in as [1;37m$username[0m[0m"
					echo "ACCESS_TOKEN=$ACCESS_TOKEN" >> .env
					echo "USERNAME=$username" >> .env
				fi
			fi
		else
			res=$(curl -s -H "Authorization: token $ACCESS_TOKEN" https://api.github.com/user)
			username=$(echo $res | ./jq-win64 -r '.login')
			if [[ $username == "null" ]]
			then
				echo "[0;31mInvalid Client Credentials ;([0m"
			else
				echo "[0;32mSuccess! Logged in as [1;37m$username[0m[0m"
			fi
		fi
		;;

	repos )
		res=$(curl -s -H "Authorization: token $ACCESS_TOKEN" https://api.github.com/user/repos)
		repos_name=$(echo $res | ./jq-win64 -r '.[].full_name')
		echo "[1;37m$repos_name[0m"
		;;

	collab )
		if [ -z $2 ]
		then
			echo "[0;31mInvalid Repository Name[0m"
			echo "$help_text"
		else
			res=$(curl -s -H "Authorization: token $ACCESS_TOKEN" https://api.github.com/repos/$USERNAME/$2/collaborators)
			collab_name=$(echo $res | ./jq-win64 -r '.[].login')
			echo "[1;37m$collab_name[0m"
		fi
		;;

	add )
		if [ -z $2 ] || [ -z $3 ]
		then
			echo "[0;31mInvalid Repository Name or Username[0m"
			echo "$help_text"
		else
			res=$(curl -s -H "Authorization: token $ACCESS_TOKEN" -X PUT https://api.github.com/repos/$USERNAME/$2/collaborators/$3)
			add_user=$(echo $res | ./jq-win64 -r '.invitee.login')
			echo "[0;32mCollaborator [1;37m$add_user[0m Invited Successfully![0m"
		fi
		;;

	remove )
		if [ -z $2 ] || [ -z $3 ]
		then
			echo "[0;31mInvalid Repository Name or Username[0m"
			echo "$help_text"
		else
			res=$(curl -s -H "Authorization: token $ACCESS_TOKEN" -X DELETE https://api.github.com/repos/$USERNAME/$2/collaborators/$3)
			echo "[0;32mCollaborator [1;37m$3[0m Removed Successfully![0m"
		fi
		;;

	--version | -v )
		echo "karvachauth VERSION: 0.0.1"
		;;

	--help | -h | * )
		echo "$help_text"
		;;

esac

