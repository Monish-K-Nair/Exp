-module(sum).
-export([main/0]).

main() -> 
    {ok, [A, B]} = io:fread("", "~d~d"),
    io:format("~p~n",[A+B]).