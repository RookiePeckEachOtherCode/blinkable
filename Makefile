run:
	sh ./run.sh
clean-run:
	rm -f ./cmd/api/api
	rm -r ./cmd/user/output
clean-log:
	find . -name "logs" | grep -v ".git" | xargs rm -rf


