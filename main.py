import dotenv
import os
import discord


def main():
    dotenv.load_dotenv()
    token = str(os.getenv("TOKEN"))
    bot = discord.Bot() #pass owner_ids and intents to this


if __name__ == "__main__":
    main()


    #client.run(token)