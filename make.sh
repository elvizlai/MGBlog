#!/bin/bash 

echo begin...

basepath=$(cd `dirname $0`; pwd)

GOPATH=$GOPATH:$basepath

cd $basepath"/src"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mblog main.go

outdir=$basepath"/output"

rm -rf $outdir

mkdir $outdir

mv $basepath"/src/mblog" $outdir

cp -r $basepath"/conf" $outdir

cp -r $basepath"/static" $outdir

cp -r $basepath"/views" $outdir

cp $basepath"/run.sh" $outdir

cp $basepath"/watermark.png" $outdir

echo end...