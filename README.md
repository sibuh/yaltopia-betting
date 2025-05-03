# yaltopia betting

# Volleyball and Cricket 

# Take this into consideration when examining the test task: 
# This code has two endpoints /vb and /crkt for volleyball and cricket matches to see the settlement of certain sample odd selection. 
# For /vb endpoint or volleyball betting sample odds are selected for only two markets WinLoss and CorrectScore
# NO WinDrawWin market in volleyball because it does not happen under regular volleyball rules.
# NO sample section for Over/Under market b/c prematch data has no adds for these market

# For cricket match i only have done sample selection for Win/Loss as Win/Draw/Win market is not usual in cricket match as is in Football matches. 
# NO sample selection for Over/Under selection for cricket match because even though over/under odds exist in prematch data cricket match does not have recored of "1st_over" and "innings_1".

# double chance markets are usual in matches with three possible outcomes (typically: Home Win, Draw, Away Win) like football. This is unusall in volleyball and cricket where three possible outcomes are unusual.Double chances can result the following outcomes.

# 1X – Home win or Draw

# 12 – Home win or Away win

# X2 – Draw or Away win. However in volleyball and 
# cricket matches where markets are dominantly Win/Loss we can have only 12 combination. So bettor can not  bet for both Home Win and Away Win.





# How to get app runing:
# option 1: Run with built in go command
# prerequest: install Golang better to have latest version like v1.23.3
# 1. open terminal or vscode terminal 
# 2. cd into yaltopia-betting folder
# 3. run go run cmd/main.go to start server

# option 2: Run with docker and docker-compose
# prerequest: need to install docker and docker compose 
# 1. open terminal or vscode terminal 
# 2. cd into yaltopia-betting folder
# 3. run docker-compose up --build if you are using python version of docker-compose  
# run docker compose up --build if you are using go version of docker compose
# 4. wait until the app starts


# To see the response of betting app endpoints 
# open browser tab and go to http://localhost:3000/vb for volleyball match and 
# http://localhost:3000/crkt for cricket match. 
# repeatedly hit the endpoints and see how response behaves 
