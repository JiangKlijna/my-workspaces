
import os
#copy /b *.ts new.ts
def files():
    for f in os.listdir('./'):
        if f.endswith('.ts'):
            yield f

def main():
    for f in files():
        os.rename(f, ('0' * (7-len(f)) + f))
        print(f)

if __name__ == '__main__':
    main()
