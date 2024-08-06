# Untruth Social

This is a project I put together for [Terrible Ideas Hackathon](https://terriblehack.com/). 

Basically, it's a little social network that only allows users to post things if they're untrue. This saves the user the guesswork of trying to find the truth hidden amongst the lies. An LLM will suggest a true post to give users some inspiration for what to write about, then when the user hits submit, if what they've written is true, the LLM will change it to be nonsense. The LLM will also be asked a third time before any post is saved to the database to make sure it's not disturbing content. 

## Running it
- Install ollama and run `ollama serve`.
- Download the model by running `ollama run llama3.1`. This will take a few minutes to download the model. You can close this once the model is downloaded.
- cd to frontend and run `yarn` and `yarn dev`. 
- cd to backend and run `go run .` or your favourite go hot reload helper. 