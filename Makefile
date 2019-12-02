BUILD_OPTS = -buildmode=pie -installsuffix cgo
GO_LDFLAGS = -tags netgo -ldflags '-linkmode external -extldflags -static -s -w'

buildDebug:
	@ go build -a -o mbex.out

build:
	@ go build -a ${GO_LDFLAGS} ${BUILD_OPTS} -o mbex