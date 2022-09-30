## Basic peer to peer cash transfer application
I built this application to understand the ACID properties of a database transaction. It can perfom operations such as account creation, transfering amount to different email addresses and return histrory of trasactions via a RESTful API.

## Local Setup

### Requirements
- Go 1.17 or higher
- Docker

## Running the application
Follow the following steps to run the application

### Clone the repository
```
git clone https://github.com/KibetBrian/PeerToPeer-Money-Transfer

```

### Install dependencies

```
cd PeerToPeer-Money-Transfer && go mod download

```

### Run the application

```
docker compose up

```