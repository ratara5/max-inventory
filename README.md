1. docker pull mariadb:10.7.4
2. docker run -d --name mariadb --port 3306:3306 --env MARIADB_ROOT_PASSWORD=rootroot mariadb:10.7.4
3. (To see into container) docker logs mariadb
4. (To see all containers running in your machine) docker ps
5. (Stop container) docker stop mariadb
6. (Restart container) docker start mariadb
7. Download and open DBeaver
    a. Click on New Database Connection icon
    b. Click on MariaDB icon
    c. Server Host: local host && Port: 3306 && Username: root && Password: rootroot
    d. Click on Connection Details
        i. Connection name: mariadb
        ii. Click on Test Connection. Download Driver if isn't downloaded yet. Should appears dialogbox: 'Connected'
    e. Click on SQL icon in dbeaver toolbar. Now is possible write a sql Script
    f. Copy squema.sql in VSC and paste it into Dbeaver 
    g. Execute Script: Click on third icon from top to bottom on the left column
    h. Right click on Databases and then click on Refresh. Should appear 'max_inventory' database.


