package com.rows;

public class Main {
    public static void main(String[] args) {
        Strings str = new Strings();

        for(int i = 0; i < 4; ++i) {
            new Thread(new Worker(i, str)).start();
        }
    }
}
