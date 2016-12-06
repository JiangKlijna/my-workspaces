
obj=main.o list.o

main: $(obj)
	gcc -o main $(obj)

main.o: list.h
	gcc -c main.c
list.o: list.h
	gcc -c list.c

clean:
	rm *.o *.exe
