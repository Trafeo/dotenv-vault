# dotenv-vault
-------------------------------------------------
By Trafeo (http://trafeo.nl)

## Functionality
 Sync an application using dotenv files with keys in vault

## Overview
This small binary written in go will take a dotenv file (.env or .env.sample) parse it, and tries to read their value from vault. The keys stored in vault should have the following format:
```
{environment}/{ENV_VARIABLE}
```

#** WARNING THIS PROJECT IS UNDER ACTIVE ALPHA DEVELOPMENT AND CANNOT BE USED IN PRODUCTION**