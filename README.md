- `scripts/run`

# How to install fast_aligner on Linux

- git clone https://github.com/clab/fast_align.git
- cd fast_align
- sudo apt-get install -y libgoogle-perftools-dev libsparsehash-dev cmake g++
- mkdir build
- cd build
- cmake ..
- make

# How to install fast_aligner on Mac

- brew install google-sparsehash
- git clone https://github.com/clab/fast_align.git
- cd fast_align
- Add `include_directories(/usr/local/Cellar/google-sparsehash/2.0.3/include)` to `CMakeLists.txt`
- mkdir build
- cd build
- cmake ..
- make
