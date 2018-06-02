#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#ifndef __CJ_BOOL
#define bool    int
#define true    1
#define false   0
#define boolstr(b) ((b) ? "true" : "false")
#endif

#define VERSION_INFO "jrun version 1.2\nauthor https://github.com/jiangKlijna\n"
#define HELP_INFO "Usage:\n\tjrun -s [source file] -m [main class] -p [parameters]\nExample:\n\tjrun -s Main.java\n\tjrun -s Main.java -p 1 2 3\n\tjrun -s *.java\n\tjrun -s Main.java Lib.java -m Main\n\tjrun -s Main.jar -m Main -p 1 2 3\n\tjrun -s Main.dex -m Main -p 1 2 3\nHelp:\n\t-s Source file\n\t-m Main Class\n\t-p Runtime parameters can be ignored\n\t-h Help\n\t-v Version\n"
#define print_help_info() {puts(HELP_INFO);exit(EXIT_FAILURE);}
#define print_version_info() {puts(VERSION_INFO);exit(EXIT_SUCCESS);}
#define print_args(argc, args) {int i = 0; for (; i < argc; i++) printf("args[%d]=%s\n", i, args[i]);}
#define print_string_array(sa) {int o = 0; for (; o < sa.size; o++) printf("%s[%d]=%s\n", sa.alias, o, sa.arr[o]);}
#define random_number(digit) (rand() % digit)

typedef struct {
    char **arr;
    size_t size;
    char *alias;
} StringArray;

// main function
int main(int argc, char **args);

// analysis parameters
void analysis_parameters(char **args, int size, StringArray **re, size_t *re_length);

// verification parameters and run
bool exec_parameters(StringArray *sas, size_t sas_size);

// run parameters
void run_parameters(StringArray *s, StringArray *m, StringArray *p);

// run java
int run_java(char *java, char *main_class, char *parameters);

// run jar
int run_jar(char *jar, char *main_class, char *parameters);

// run dex
int run_dex(char *dex, char *main_class, char *parameters);

// match suffix name
bool is_match_suffix(char *file, char *suffix);

// string array join
char *join(char **strs, size_t strs_size, char *sep);

// get random cache dir size=18
char *random_cache_dir();

// clean cache dir and free size=18
void clean_cache_dir(char *);

// string.lastIndexOf
int lastIndexOf(char *str, char e);

int main(int argc, char **args) {
//    print_args(argc, args);
    if (argc == 1) print_help_info();
    if (strcmp("-h", args[1]) == 0) print_help_info();
    if (strcmp("-v", args[1]) == 0) print_version_info();
    size_t sas_size;
    StringArray *sas;
    analysis_parameters(args + 1, argc - 1, &sas, &sas_size);
    sas[0].alias = args[0];

    bool isValid = exec_parameters(sas, sas_size);
    free(sas);
    if (!isValid) print_help_info();
    exit(EXIT_SUCCESS);
}

void analysis_parameters(char **args, int size, StringArray **re, size_t *re_length) {
    size_t i, j = 0;

    size_t re_size = 1;
    for (i = 0; i < size; i++) {
        if (args[i][0] == '-') re_size++;
    }
    StringArray *result = calloc(re_size, sizeof(StringArray));
    int result_i = 0;

    char *alias = NULL;
    for (i = 0; i < size; ++i) {
        if (args[i][0] == '-') {
            StringArray sa = {args + i - j, j, alias};
            result[result_i++] = sa;
            j = 0;
            alias = args[i];
        } else {
            ++j;
        }
    }
    StringArray sa = {args + i - j, j, alias};
    result[result_i] = sa;

    *re = result;
    *re_length = re_size;
}

bool exec_parameters(StringArray *sas, size_t sas_size) {
    int i = 1;
    StringArray *source = NULL;
    StringArray *main = NULL;
    StringArray *params = NULL;

    for (; i < sas_size; ++i) {
        StringArray sa = *(sas + i);
        if (strcmp(sa.alias, "-s") == 0) {
            source = sas + i;
        } else if (strcmp(sa.alias, "-m") == 0) {
            main = sas + i;
        } else if (strcmp(sa.alias, "-p") == 0) {
            params = sas + i;
        }
    }
    if (source == NULL || source->size == 0) {
        if (sas->size == 0) return false;
        source = sas;
    }
    if (is_match_suffix(source->arr[0], "java")) {
        int j = 1;
        for (; j < source->size; ++j) {
            if (!is_match_suffix(source->arr[j], "java")) return false;
        }
    } else if (is_match_suffix(source->arr[0], "jar")) {
        if (source->size > 1) return false;
    } else if (is_match_suffix(source->arr[0], "dex")) {
        if (source->size > 1) return false;
    } else return false;
    if (main != NULL && main->size > 1) return false;

    run_parameters(source, main, params);
    return true;
}

void run_parameters(StringArray *s, StringArray *m, StringArray *p) {
    char *parameters = (p == NULL || p->size == 0) ? NULL : join(p->arr, p->size, " ");

    int point_index = lastIndexOf(s->arr[0], '.');
    int slash_index = lastIndexOf(s->arr[0], '/') + 1;

    int filename_size = point_index - slash_index;
    char filename[filename_size];
    memcpy(filename, s->arr[0] + slash_index, filename_size);
    filename[filename_size] = '\0';

    char *filetype = s->arr[0] + point_index + 1;
    char *main = (m == NULL || m->size == 0) ? filename : m->arr[0];

    if (strcmp(filetype, "java") == 0) {
        char *source = join(s->arr, s->size, " ");
        run_java(source, main, parameters);
        free(source);
    } else if (strcmp(filetype, "jar") == 0) {
        run_jar(s->arr[0], main, parameters);
    } else if (strcmp(filetype, "dex") == 0) {
        run_dex(s->arr[0], main, parameters);
    }
    free(parameters);
}

int run_java(char *java, char *main_class, char *parameters) {
    int status;
    char *cache_dir = random_cache_dir();

    char *ecj_strs[4] = {"ecj", java, "-d", cache_dir};
    char *ecj_cmd = join(ecj_strs, 4, " ");
    status = system(ecj_cmd);
    free(ecj_cmd);
    if (status != 0) {
        clean_cache_dir(cache_dir);
        return status;
    }
    char *dx_strs[3] = {"cd ", cache_dir," && dx --dex --output=run.dex *.class"};
    char *dx_cmd = join(dx_strs, 3, "");
    status = system(dx_cmd);
    free(dx_cmd);
    if (status != 0) {
        clean_cache_dir(cache_dir);
        return status;
    }

    char *dex_strs[2] = {cache_dir, "run.dex"};
    char *dex_name = join(dex_strs, 2, "");
    status = run_dex(dex_name, main_class, parameters);
    free(dex_name);

    clean_cache_dir(cache_dir);
    return status;
}

int run_jar(char *jar, char *main_class, char *parameters) {
    int status;
    char *cache_dir = random_cache_dir();

    char *dx_strs[6] = {"mkdir ", cache_dir," && dx --dex --output=", cache_dir, "run.dex ", jar};
    char *dx_cmd = join(dx_strs, 6, "");
    status = system(dx_cmd);
    free(dx_cmd);
    if (status != 0) {
        clean_cache_dir(cache_dir);
        return status;
    }

    char *dex_strs[2] = {cache_dir, "run.dex"};
    char *dex_name = join(dex_strs, 2, "");
    status = run_dex(dex_name, main_class, parameters);
    free(dex_name);

    clean_cache_dir(cache_dir);
    return status;
}

int run_dex(char *dex, char *main_class, char *parameters) {
    char *strs[4] = {"dalvikvm -cp", dex, main_class, parameters};
    char *cmd = join(strs, parameters ? 4 : 3, " ");
    int status = system(cmd);
    free(cmd);
    return status;
}

bool is_match_suffix(char *file, char *suffix) {
    size_t i, j;
    size_t file_length = strlen(file);
    size_t suffix_length = strlen(suffix);
    for (i = file_length - 1; i >= 0; --i) {
        if (file[i] == '.') {
            i++;
            break;
        }
    }
//    for (i = 0; i < file_length; ++i) {
//        if (file[i] == '.') {
//            i++;
//            break;
//        }
//    }
    for (j = 0; j < suffix_length; ++j, ++i) {
        if (suffix[j] != file[i]) return false;
    }
    if (i != file_length) return false;
    return true;
}

char *join(char **strs, size_t strs_size, char *sep) {
    size_t i, length = (strs_size - 1) * strlen(sep);
    for (i = 0; i < strs_size; ++i) {
        length += strlen(strs[i]);
    }
    char *str = calloc(length, sizeof(char));
    for (i = 0; i < strs_size - 1; ++i) {
        strcat(str, strs[i]);
        strcat(str, sep);
    }
    strcat(str, strs[i]);
    return str;
}

int lastIndexOf(char *str, char e) {
    int i = (int) strlen(str);
    for (; i >= 0; --i) {
        if (str[i] == e) return i;
    }
    return -1;
}

char *random_cache_dir() {
    unsigned t = (unsigned) time(0);
    srand(t);
    unsigned r = (unsigned) random_number(999);

    char *cache_dir = calloc(18, sizeof(char));
    sprintf(cache_dir, ".out%d%d/", t, r);
    return cache_dir;
}

void clean_cache_dir(char *dir) {
    char cmd[7 + 18];
    sprintf(cmd, "rm -rf %s", dir);
    system(cmd);
    free(dir);
}
