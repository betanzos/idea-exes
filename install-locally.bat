:: Remove output directory
rd /s /q target

:: Create output directory for EXE files
mkdir target

:: Create executable for IntelliJ IDEA
cd main
go build -ldflags -H=windowsgui -o "../target/custom-idea.exe"
cd ..
copy hidebat.vbs target\