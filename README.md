# code_work_telegramm
Для работы

export GO111MODULE=on

go get -u github.com/gin-gonic/gin

go get github.com/gin-contrib/cors


  //api/instruments/list
// https://habr.com/ru/post/495324/

// https://iss.moex.com/iss/history/engines/stock/markets/shares/boardgroups/57/securities.jsonp?iss.meta=off&iss.json=extended&lang=ru&security_collection=3&date=2022-02-08&start=200&limit=100&sort_column=VALUE&sort_order=des

 // короткие имена акций
// https://iss.moex.com/iss/engines/stock/markets/shares/boards/TQBR/securities.json?iss.meta=off&iss.only=securities&securities.columns=SECID,SECNAME

//Узнавать текущую цену для конкретной ценной бумаги
//http://iss.moex.com/iss/engines/stock/markets/shares/boards/TQBR/securities.json?iss.meta=off&iss.only=securities&securities.columns=SECID,PREVADMITTEDQUOTE
 