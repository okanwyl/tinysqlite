BUILD_TYPE ?= debug

SRCDIR = src
INCDIR = include
BUILDDIR = build

CC = clang
CFLAGS_debug = -g -O0
CFLAGS_release = -O3
CFLAGS_common = -Wall -I$(INCDIR)

CFLAGS = $(CFLAGS_common) $(CFLAGS_$(BUILD_TYPE))

# Linker
LDFLAGS =

SRC = $(wildcard $(SRCDIR)/*.c)
OBJ = $(SRC:$(SRCDIR)/%.c=$(BUILDDIR)/%.o)

all: $(BUILDDIR)/tinysqlite

$(BUILDDIR)/tinysqlite: $(OBJ)
	$(CC) $(OBJ) -o $@ $(LDFLAGS)

$(BUILDDIR)/%.o: $(SRCDIR)/%.c
	mkdir -p $(BUILDDIR)
	$(CC) -c $< -o $@ $(CFLAGS)

clean:
	rm -rf $(BUILDDIR)

