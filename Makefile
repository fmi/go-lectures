# A Make file, that conforms with the web site's expectations. Running 'make'
# has to create a directory called 'compiled-lectures' that will be linked by
# the site's deploy process.
#
# "Compiling" the Go lectures does not mean much. We just copy the files to their
# new location.

TEMPDIR := $(shell mktemp -d)
COMPILED = compiled-lectures

all:
	rm -rf $(COMPILED)
	cp -a ./* "$(TEMPDIR)/"
	mv $(TEMPDIR) $(COMPILED)

clean:
	rm -rf $(COMPILED)
