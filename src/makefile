BINDIR=../bin

OBJS = permute \
			 gsort

GOFLAGS =

.SUFFIXES: .go

all: bindir $(OBJS)

bindir:
	mkdir -p $(BINDIR)
.go:
	go build -o $(BINDIR)/$* $(GOFLAGS) $<
