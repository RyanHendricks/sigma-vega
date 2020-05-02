#!/bin/bash
docker run -it --rm -v "$(pwd)/.github":/usr/local/src/your-app ferrarimarco/github-changelog-generator -u RyanHendricks -p sigma-vega --token $GHTOKEN
