#include "hex.h"
#include <cmath>
#include <fstream>
#include <set>
using namespace std;

const string Hex::BASE(
        "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ");

Hex::Hex(const unsigned bit) :
    BASE_CHAR(Hex::BASE.substr(0, bit)), BASE_NUM(bit) {
}
Hex::Hex(const string &base_char) :
    BASE_CHAR(base_char), BASE_NUM(base_char.size()) {
}

long Hex::x_ten(const string &x) {
    long num = 0;
    for (int i = x.size() - 1; i >= 0; i--) {
        num += query(x[i]) * pow(BASE_NUM, x.size() - 1 - i);
    }
    return num;
}

string Hex::ten_x(long num) {
    string x = "";
    while (num > 0) {
        int a = num % BASE_NUM;
        x = BASE_CHAR[a] + x;
        num /= BASE_NUM;
    }
    return x;
}

int Hex::query(char c) {
    for (int i = 0, n = BASE_CHAR.size(); i < n; i++) {
        if (c == BASE_CHAR[i]) {
            return i;
        }
    }
    return -1;
}

void Hex::encrypFile(const string &srcf, const string &descf){
    ifstream in(srcf);
    if (!in.is_open()){
        return;
    }
    ofstream out(descf);
    if (!out.is_open()){
        return;
    }
    while (!in.eof()){
        out << ten_x(in.get()) + ",";
    }
    in.close();
    out.close();
}

void Hex::decrypFile(const string &srcf, const string &descf){
    char temp[64];

    ifstream in(srcf);
    if (!in.is_open()){
        return;
    }
    ofstream out(descf);
    if (!out.is_open()){
        return;
    }
    while (!in.eof()){
        in.getline(temp, 64, ',');
        out << (char)x_ten(temp);
    }
    in.close();
    out.close();
}

//string Hex::encrypStr(const string &srcs){
//    string str;
//    for(string::iterator i = srcs.begin(); i != srcs.end(); i++){
//        str.append(ten_x(*i));
//    }
//    return str;
//}

//string Hex::decrypStr(const string &srcs){
//    string str;
//    for(string::iterator i = srcs.begin(); i != srcs.end(); i++){
//        str.append(x_ten(*i));
//    }
//    return str;
//}

bool Hex::testKey(const unsigned &bit){
    return bit > 1 && bit <= 36;
}

bool Hex::testKey(const string &base_char){
    unsigned size = base_char.size();
    if(size < 6){
        return false;
    }
    set<int> data;
    for(unsigned i = 0; i < size; i++){
        int c = base_char[i];
        if(c<0 || c>255)
            return false;
        data.insert(i);
    }
    return data.size() == size;
}

