****    aggreGator    ****

    This program aggregates rss feeds.
Using the cli you can save rss feeds and 
download them into a local database.

Requirements:

* postgres
* Go (golang)

Getting Started:
    
    In order to install download the files and 
open up a shell. Once in the shell run go install in 
the downloaded folder.

Config:

In your home file you will need a config file.
* .gatorconfig.json
    It contains 2 things: 
    * The connection string for the postgres database
    * The last user that was logged in

Initially just put in the connection string.
eg. "postgres://postgres:postgres@localhost:5432/gator"
current_user_name value can be left empty (i.e. "")

Example config
{
"db_url":
    "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
"current_user_name":
    "kahya"
}

Once installed you have several commands.
Used in the form of "gator <command name>"

* register <username>
    Allows you to register users
* login <username>
    Allows you to switch users
* users
    Lists all users
* follow <Feed URL>
    Follows an RSS feed that is already existing
* unfollow <Feed URL>
    Unfollows specified RSS feed
* following
    Lists all feeds followed by current user
* feeds
    Lists all feeds currently followed by any user
* addfeed <Feed name> <Feed URL>
    Adds a feed and follows it for the current user
* agg <Refresh timer>
    Runs continuously and updates feeds
    Timer should be in the form of 1m (1 minute), 2h (2 hours)
    <<<< Caution with this command, too short of a 
    timer may overload resources>>>>
* browse <optional limit>
    Shows you the most recent posts.
    If limit is not specified it defaults to 2.
    Limit must be in integer form (1, 2, 3, 4, etc)
    


