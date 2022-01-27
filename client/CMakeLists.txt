cmake_minimum_required(VERSION 3.20)
project(qt_client)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)
set(CMAKE_AUTOUIC ON)

find_package(Qt6 COMPONENTS Core)
find_package(Qt6 COMPONENTS Gui)
find_package(Qt6 COMPONENTS Quick)
find_package(Qt6 COMPONENTS WebEngineQuick)

add_executable(main main.cc main.qml main.qrc)

target_link_libraries(main PUBLIC Qt::Core Qt::Gui Qt::Quick Qt::WebEngineQuick)
