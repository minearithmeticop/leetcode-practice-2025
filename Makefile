.PHONY: test test-all test-1480 fmt lint clean help

## แสดงวิธีใช้
help:
	@echo.
	@echo วิธีใช้งาน:
	@echo   make ^<target^>
	@echo.
	@echo คำสั่งที่ใช้งานได้:
	@echo   test-all     รันเทสทั้งหมด
	@echo   test-1480    รันเทสเฉพาะปัญหา 1480
	@echo   fmt          จัดรูปแบบโค้ดอัตโนมัติ
	@echo   lint         ตรวจสอบคุณภาพโค้ด
	@echo   clean        ล้างไฟล์ build
	@echo   help         แสดงวิธีใช้

## รันเทสทั้งหมด
test-all:
	@echo [กำลังรันเทสทั้งหมด...]
	@go test -v ./...

## รันเทสเฉพาะปัญหา 1480
test-1480:
	@echo [กำลังรันเทสปัญหา 1480...]
	@cd problems/problem-1480 && go test -v

## จัดรูปแบบโค้ดอัตโนมัติ
fmt:
	@echo [กำลังจัดรูปแบบโค้ด...]
	@go fmt ./...

## ตรวจสอบคุณภาพโค้ด
lint:
	@echo [กำลังตรวจสอบคุณภาพโค้ด...]
	@if not exist "$(GOPATH)\bin\golangci-lint.exe" @echo โปรดติดตั้ง golangci-lint ก่อน: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && exit /b 1
	@golangci-lint run ./...

## ล้างไฟล์ build
clean:
	@echo [กำลังล้างไฟล์ที่ไม่จำเป็น...]
	@go clean

## แสดงวิธีใช้
default: help
