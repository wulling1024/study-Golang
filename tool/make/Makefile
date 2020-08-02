result.txt: source.txt
	cp source.txt result.txt
source.txt:
	echo "This is the source" > source.txt

new: file1 file2 file3

file1:
	touch file1.txt
file2:
	touch file2.txt
file3:
	touch file3.txt

del: dfile1 dfile2 dfile3

dfile1:
	rm file1.txt
dfile2:
	rm file2.txt
dfile3:
	rm file3.txt
