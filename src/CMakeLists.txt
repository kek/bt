cmake_minimum_required(VERSION 3.21)

include(${CMAKE_CURRENT_SOURCE_DIR}/cmake/prelude.cmake)
project(simpleble_examples)
include(${CMAKE_CURRENT_SOURCE_DIR}/cmake/epilogue.cmake)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_EXTENSIONS OFF)

option(SIMPLEBLE_LOCAL "Use local SimpleBLE" ON)

if (SIMPLEBLE_LOCAL)
    add_subdirectory(${CMAKE_CURRENT_LIST_DIR}/../../SimpleBLE/simpleble ${CMAKE_BINARY_DIR}/simpleble)
else()
    find_package(simpleble CONFIG REQUIRED)
endif()

# C
add_executable(example_connect_c connect.c)
target_link_libraries(example_connect_c simpleble::simpleble-c)

add_executable(example_notify_c notify.c)
target_link_libraries(example_notify_c simpleble::simpleble-c)

add_executable(example_scan_c scan.c)
target_link_libraries(example_scan_c simpleble::simpleble-c)

add_library(bluetooth bluetooth.c)
target_link_libraries(bluetooth simpleble::simpleble-c)

add_library(bluetooth_example bluetooth_example.c)
target_link_libraries(bluetooth_example simpleble::simpleble-c)
