#!/bin/sh

wget https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.4/wkhtmltox-0.12.4_linux-generic-amd64.tar.xz

tar -xf wkhtmltox-0.12.4_linux-generic-amd64.tar.xz

cp wkhtmltox/bin/wkhtmltoimage /usr/bin/

cp wkhtmltox/bin/wkhtmltopdf /usr/bin/