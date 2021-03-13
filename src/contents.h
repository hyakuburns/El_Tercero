#ifndef CONTENTS_H_
#define CONTENTS_H_

char* RDME = "# %s";
char* CCONT =
    "#include <stdio.h>\n\n"
    "int main(int argc, char** argv){\n"
    "  printf(\"Eternitatem cogita.\\n\");\n"
    "}";
char* CMES =
    "project(\'%s\', \'c\', default_options : [\'warning_level=3\', "
    "\'c_std=c99\'])\n\n"

    "add_global_arguments(language : \'c\')\n"
    "c = meson.get_compiler(\'c\')\n\n"

    "#Sample built-in libs\n"
    "#mathl = c.find_library(\'m\', required : true)\n"
    "#runtimel = c.find_library(\'rt\', required : true)\n"
    "#dlib = c.find_library(\'dl\', required : true)\n\n"

    "#Sample dependencies\n"
    "#depz = [dependency(\'raylib\'), dependency(\'GL\'), "
    "dependency(\'X11\'), dependency(\'threads\')]\n\n"

    "#include dirs\n"
    "#libz = include_directories(\'lib\')\n\n"

    "#main source\n"
    "srcs = [\'src/main.c\']\n\n"

    "executable(\'main\', srcs)";

char* CXXCONT =
    "#include <iostream>\n\n"
    "int main(int argc, char** argv){\n"
    "  printf(\"Extra lutum pedes habes.\\n\");\n"
    "}";
char* CXXMES =
    "project(\'%s\', \'cpp\', default_options : [\'warning_level=3\', "
    "\'cpp_std=c++17\'])\n\n"

    "add_global_arguments(language : \'cpp\')\n"
    "cpp = meson.get_compiler(\'cpp\')\n\n"

    "#Sample built-in libs\n"
    "#mathl = cpp.find_library(\'m\', required : true)\n"
    "#runtimel = cpp.find_library(\'rt\', required : true)\n"
    "#dlib = cpp.find_library(\'dl\', required : true)\n\n"

    "#Sample dependencies\n"
    "#depz = [dependency(\'raylib\'), dependency(\'GL\'), "
    "dependency(\'X11\'), dependency(\'threads\')]\n\n"

    "#include dirs\n"
    "#libz = include_directories(\'lib\')\n\n"

    "#main source\n"
    "srcs = [\'src/main.cxx\']\n\n"

    "executable(\'main\', srcs)";
char* GATR =
    "# Sources\n"
    "*.c     text diff=c\n"
    "*.cc    text diff=cpp\n"
    "*.cxx   text diff=cpp\n"
    "*.cpp   text diff=cpp\n"
    "*.c++   text diff=cpp\n"
    "*.hpp   text diff=cpp\n"
    "*.h     text diff=c\n"
    "*.h++   text diff=cpp\n"
    "*.hh    text diff=cpp\n"
    "# Compiled Object files\n"
    "*.slo   filter=lfs diff=lfs merge=lfs -binary\n"
    "*.lo    filter=lfs diff=lfs merge=lfs -binary\n"
    "*.o     filter=lfs diff=lfs merge=lfs -binary\n"
    "*.obj   filter=lfs diff=lfs merge=lfs -binary\n"
    "# Precompiled Headers\n"
    "*.gch   filter=lfs diff=lfs merge=lfs -binary\n"
    "*.pch   filter=lfs diff=lfs merge=lfs -binary\n"
    "# Compiled Dynamic libraries\n"
    "*.so    filter=lfs diff=lfs merge=lfs -binary\n"
    "*.dylib filter=lfs diff=lfs merge=lfs -binary\n"
    "*.dll   filter=lfs diff=lfs merge=lfs -binary\n"
    "# Compiled Static libraries\n"
    "*.lai   filter=lfs diff=lfs merge=lfs -binary\n"
    "*.la    filter=lfs diff=lfs merge=lfs -binary\n"
    "*.a     filter=lfs diff=lfs merge=lfs -binary\n"
    "*.lib   filter=lfs diff=lfs merge=lfs -binary\n"
    "# Executables\n"
    "*.exe   filter=lfs diff=lfs merge=lfs -binary\n"
    "*.out   filter=lfs diff=lfs merge=lfs -binary\n"
    "*.app   filter=lfs diff=lfs merge=lfs -binary\n"
    "*      text=auto\n"
    "*.clip filter=lfs diff=lfs merge=lfs -binary\n"
    "*.png  filter=lfs diff=lfs merge=lfs -binary\n"
    "*.jpg  filter=lfs diff=lfs merge=lfs -binary\n"
    "xxx    filter=lfs diff=lfs merge=lfs -binary\n"
    "*.out  filter=lfs diff=lfs merge=lfs -binary\n";
char* GIGN =
    "#Emacs stuff\n"
    "*~\n"
    ".#*\n"
    "# ---> C\n"
    "# Prerequisites\n"
    "*.d\n\n"

    "# Object files\n"
    "*.o\n"
    "*.kov\n"
    "*.obj\n"
    "*.elf\n\n"

    "# Linker output\n"
    "*.ilk\n"
    "*.map\n"
    "*.exp\n\n"

    "# Precompiled Headers\n"
    "*.gch\n"
    "*.pch\n\n"

    "# Libraries\n"
    "*.lib\n"
    "*.a\n"
    "*.la\n"
    "*.lo\n\n"

    "# Shared objects (inc. Windows DLLs)\n"
    "*.dll\n"
    "*.so\n"
    "*.so.*\n"
    "*.dylib\n\n"

    "# Executables\n"
    "*.exe\n"
    "*.out\n"
    "*.app\n"
    "*.i*86\n"
    "*.x86_64\n"
    "*.hex\n"
    "xxx\n\n"

    "# Debug files\n"
    "*.dSYM/\n"
    "*.su\n"
    "*.idb\n"
    "*.pdb\n\n"

    "# Kernel Module Compile Results\n"
    "*.mod*\n"
    "*.cmd\n"
    ".tmp_versions/\n"
    "modules.order\n"
    "Module.symvers\n"
    "Mkfile.old\n"
    "dkms.conf\n\n"

    "# ---> C++\n\n"

    "# Compiled Object files\n"
    "*.slo\n\n"

    "# Fortran module files\n"
    "*.mod\n"
    "*.smod\n\n"

    "# Compiled Static libraries\n"
    "*.lai\n"
    "*.a\n\n"

    "# Executables\n"
    "*.exe\n"
    "*.out\n"
    "*.app\n\n"

    "# ---> Rust\n\n"

    "# Generated by Cargo"
    "# will have compiled files and executables\n"
    "debug/\n"
    "target/\n\n"

    "# Remove Cargo.lock from gitignore if creating an executable, leave it "
    "for libraries\n"
    "# More information here"
    " https://doc.rust-lang.org/cargo/guide/cargo-toml-vs-cargo-lock.html\n"
    "Cargo.lock\n\n"

    "# These are backup files generated by rustfmt\n"
    "**/*.rs.bk\n\n"

    "#RoboDeps\n"
    "robodep/\n\n"

    "# Meson stuff\n"
    ".cache/\n"
    "build/\n"
    "debugxxx\n"
    "test/\n"
    "%s";

#endif /* ifndef CONTENTS_H_ */
