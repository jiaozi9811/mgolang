#!/bin/env bash

sha256string(){
    string=$1
    sum=`echo ${string}|sha256sum|awk '{print $1}'`
    echo $sum
    #return $sum
}

sha256file(){
    hashfile=$1
    sum=`sha256sum ${hashfile}|awk '{print $1}'`
    echo $sum
    #return $sum
}

c=$(sha256string dsfsd)
echo $c

c=$(sha256file ./anaconda-ks.cfg)
echo $c
