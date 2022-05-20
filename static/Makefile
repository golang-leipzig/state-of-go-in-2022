CHANGES_PNG := \
	changes1.14.png \
	changes1.15.png \
	changes1.16.png \
	changes1.17.png \
	changes1.18.png \

changes-combined.png: $(CHANGES_PNG)
	convert $^ -append $@

.PHONY: clean
clean:
	rm -f changes-combined.png
