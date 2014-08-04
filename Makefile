.PHONY: clean

default: build

clean:
	@go clean
	@rm -rf pkg

build:
	@gox -os "darwin linux windows" -output "pkg/{{.OS}}_{{.Arch}}/{{.Dir}}"

install:
	@go install
