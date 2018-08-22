# softuni-camp-miners-game

## Summary 

This is just a simple game I made during a lecture, where
I spoke about the **Theory of Concept** in Programming, 
at the **SoftUni Camp** (__Balchick, 2018__)

It was a really funny experience, because the communication with the auditory of students was high,
we had a laugh, we had fun, and we implemented a simple game.

I provided them with a lot of freedom - their task was to choose a random technology, on which we would
create a simple game.
I also provided them with 1 rule, which was supposed to be a great adventure for me - their task was to
choose a technology, I've never seen, wrote or configured before.

Due to my extreme curiosity, I've introduced myself to a great amount of technologies. Of course, I've not
reached the heights of development in those technologies, but I really wanted a technology I've never seen
before.

There were 2 particularly popular technologies, which I knew would be my backup choice - **Go** & **R**.
It was a great ordeal for me to **NOT** touch those technologies before the scheduled lecture, but I somehow did it.

The students choose **Go Lang**, for one reason or another, so the game, which's source code you would see is
purely written on **Go Lang**. 

## Lecture Specification

The main topic of the lecture was: Do not write a particular technology... Write code! 
Write high-quality code! Write optimized code! Write conceptually-right code! 

Most of the concepts of programming are shared among the different technologies
and programming languages. There is a moment to the professional life of each good developer, 
at which moment he / she / it reaches a certain summit. And at that summit, he / she / it discovers
the ability to develop software on any technology, with just a simple effort.

That is what it means to be a real programmer.

## Game Specification

The game is very simple. You start with 50 currency (the currency is not specified) 
The game's models are **Miners** and **Minerals**. The miners mine minerals on a
random basis. Those Minerals are collected, and stored. Later you can sell them.
Each resource worths different amount of currency. 

The main idea is just to mine minerals. The game is extremely simple, it was just a 
funny idea.

There is currently only one Miner - **BasicMiner**, with the idea that it can be extended in the future.
There is currently 6 Minerals with different currencies - **Dust**, **Coal**, **Copper**, **Iron**, **Silver**, **Gold**.

There are several commands:

Miner - Buys a BasicMiner (costs 10 currency)
Harvest - All Miners mine minerals (costs 5 currency per miner)
Sell - Sells all mined minerals
Status - Brings detailed information about the count of miners, the mined resources and the current budget.
Quit - Quits the game

## Some Disclaimers

The code is not that high-quality, and probably breaks some of the principles of high-quality code in Go Lang.
However, it is quite conceptually correct. By following the OOP principles of most languages, this game
was implemented quite similar to a normal project on C# or Java for example.
 