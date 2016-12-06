#ifndef HEX_H
#define HEX_H

#include <iostream>

class Hex {
public:
    Hex(const unsigned bit = 10);
    Hex(const std::string &base_char);

    long x_ten(const std::string &x);
    std::string ten_x(long num);

    void encrypFile(const std::string &srcf, const std::string &descf);
    void decrypFile(const std::string &srcf, const std::string &descf);

    std::string encrypStr(const std::string &srcs);
    std::string decrypStr(const std::string &srcs);

    static bool testKey(const unsigned &bit);
    static bool testKey(const std::string &base_char);


private:
    const static std::string BASE;
    const std::string BASE_CHAR;
    const unsigned BASE_NUM;

    int query(char c);
};

#endif // HEX_H
