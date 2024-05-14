cmake_minimum_required(VERSION 3.14)
project(cityproto LANGUAGES CXX)

# global settings
# activate intellisense
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_FLAGS " ${CMAKE_CXX_FLAGS_INIT} -fPIC")

# include cmake
set(CMAKE_MODULE_PATH "${PROJECT_SOURCE_DIR}/cmake/" ${CMAKE_MODULE_PATH})
include(BuildType)

# set flag
set(ABSL_ENABLE_INSTALL ON)
include(FetchContent)
FetchContent_Declare(
  grpc
  GIT_REPOSITORY https://github.com/grpc/grpc.git
  GIT_TAG v1.54.3
  GIT_SUBMODULES_RECURSE TRUE
)
set(FETCHCONTENT_QUIET OFF)
FetchContent_MakeAvailable(grpc)

# find all .h and .cc files in cpp/ directory
file(GLOB_RECURSE HEADER_FILES "${CMAKE_CURRENT_SOURCE_DIR}/cpp/city/**/*.h")
file(GLOB_RECURSE SOURCE_FILES "${CMAKE_CURRENT_SOURCE_DIR}/cpp/city/**/*.cc")
message(DEBUG "HEADER_FILES: ${HEADER_FILES}")
message(DEBUG "SOURCE_FILES: ${SOURCE_FILES}")

add_library(cityproto ${HEADER_FILES} ${SOURCE_FILES})
message(DEBUG "INCLUDE: ${CMAKE_CURRENT_BINARY_DIR}/cpp")
target_include_directories(cityproto PUBLIC "${CMAKE_CURRENT_SOURCE_DIR}/cpp")
target_link_libraries(
  cityproto
  PUBLIC grpc++
  PUBLIC protobuf::libprotobuf
)

set_target_properties(cityproto PROPERTIES
  CXX_STANDARD 17
  CXX_STANDARD_REQUIRED YES
)
