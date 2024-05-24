# Configuration

In the first part of this series, we added a library to autoload our environment, either
from the environment or from the env files (.env, .env.local, etc). 

Now, since we have our environment variables loaded, we use another library 
to parse our env variables into our own defined configuration struct.

The configuration struct lives inside `internal/`, since it's actually supposed to be
private and used only by our `cmd/basic-service` command. 