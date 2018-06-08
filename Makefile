build:
	docker build -t vothanhkiet/noop:latest .

upload:
	docker push vothanhkiet/noop:latest

clean:
	rm -rf debug
