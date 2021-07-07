## Install PostgreSQL on Ubuntu 

1. `Run the following commands to install PostgreSQL server on Ubuntu`  

#### Step 1 Install PostgreSQL
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```  
* We’re also installing the PostgreSQL contrib package that provides several additional features for the PostgreSQL database system.   
* Once the installation is completed, the PostgreSQL service will start automatically.     
*  Use the psql tool to verify the installation by connecting to the PostgreSQL database server and printing its version.   
```bash
sudo -u postgres psql -c "SELECT version();"
```  
* successful installation verify the postgresql service  
```bash
sudo systemctl status postgresql 
```

#### Step 2  Connection To PostgreSQL  
* Now, establish a connection with the newly installed Postgres database server.    
*  First switch to the system’s postgres user account:  
```bahs
sudo su - postgres 
``` 
* then type “psql” to get the postgress prompt:  
```bash
psql
```
* Once you are connected to PostgreSQL and you can see the connection information’s details, use the following command:   

** Note remove one slash because .md file add one extra slash.   
```bash
postgres=# \conninfo
```  

####  Step 3 – Secure PostgreSQL     

* PostgreSQL installer creates a user “postgres” on your system. Default this user is not protected.  
* First, create a password for “postgres” user account by running the following command.   
```bash
sudo passwd postgres 
```
* Next, switch to the “postgres” account Then switch to the Postgres system account and create a secure and strong password    
  for PostgreSQL administrative database user/role as follows.  
  
```bash
su - postgres 
psql -c "ALTER USER postgres WITH PASSWORD 'secure_password_here';" 
exit 
```
* Restart the service to apply security changes.
```bash
sudo systemctl restart postgresql 
```
#### Step 4 – Install pgAdmin
* First, import the repository signing GPG key and add the pgAdmin4 PPA to your system using the following commands.  
```bash
curl https://www.pgadmin.org/static/packages_pgadmin_org.pub | sudo apt-key add -
sudo sh -c 'echo "deb https://ftp.postgresql.org/pub/pgadmin/pgadmin4/apt/focal pgadmin4 main" > /etc/apt/sources.list.d/pgadmin4.list' 
``` 

* After adding the PPA, update the Apt cache and install the pgAdmin4 package on your system.  
```bash
sudo apt update
sudo apt install pgadmin4 
```

```bash
sudo /usr/pgadmin4/bin/setup-web.sh
```
* The above script will prompt you to create a user to access the web interface. Input an email address and password when prompted.  
* Say “y” for other confirmation asked by the script.  

#### Step 5 – use pgAdmin in web browser
* Enter url in web browser `http://localhost/pgadmin4`  
* Enter email and password which is created in step 4.  
* Click `Add New Server`  
* Click connection `Host name/ address` : `127.0.0.1`  
* Click connection Enter user name and password for step 3  

### Key Points 
* Database access permissions within PostgreSQL are handled with the concept of roles.   
* A role can represent a database user or a group of database users.  
