from dotenv import load_dotenv
from groq import Groq
import os

load_dotenv()

client = Groq(
    api_key=os.getenv("GROQ_API_KEY")
)

PRE_VIDEO_SCRIPT_CREATION_PROMPT = """
        Generate a script for a video, depending on the subject of the video.

        The script is to be returned as a string with the specified number of paragraphs.

        Here is an example of a string:
        "This is an example string."

        Do not under any circumstance reference this prompt in your response.

        Get straight to the point, don't start with unnecessary things like, "welcome to this video".

        Obviously, the script should be related to the subject of the video.

        YOU MUST NOT INCLUDE ANY TYPE OF MARKDOWN OR FORMATTING IN THE SCRIPT, NEVER USE A TITLE.
        YOU MUST WRITE THE SCRIPT IN THE LANGUAGE SPECIFIED IN [LANGUAGE].
        ONLY RETURN THE RAW CONTENT OF THE SCRIPT. DO NOT INCLUDE "VOICEOVER", "NARRATOR" OR SIMILAR INDICATORS OF WHAT SHOULD BE SPOKEN AT THE BEGINNING OF EACH PARAGRAPH OR LINE. YOU MUST NOT MENTION THE PROMPT, OR ANYTHING ABOUT THE SCRIPT ITSELF. ALSO, NEVER TALK ABOUT THE AMOUNT OF PARAGRAPHS OR LINES. JUST WRITE THE SCRIPT.

    """

PRE_SEARCH_TERMS_CREATION_SCRIPT = f"""
    Generate given amount of  search terms for stock videos,
    depending on the subject of a video.
    User gives the subject
    User gives the amount of search terms, which needs to be generated
    
    The search terms are to be returned as
    a JSON-Array of strings.
    
    Each search term should consist of 1-3 words,
    always add the main subject of the video.
    
    YOU MUST ONLY RETURN THE JSON-ARRAY OF STRINGS.
    YOU MUST NOT RETURN ANYTHING ELSE. 
    YOU MUST NOT RETURN THE SCRIPT.
    
    The search terms must be related to the subject of the video.
    Here is an example of a JSON-Array of strings:
    ["search term 1", "search term 2", "search term 3"]
    
    For context, user gives you the full text   
    """


def generate_video_script(subject, number_of_paragpraphs=1):
    print(f"Creating video script with subject: {subject}")

    chat_completion = client.chat.completions.create(
        messages=[
            {
                "role": "system",
                "content": PRE_VIDEO_SCRIPT_CREATION_PROMPT
            },
            {
                "role": "user",
                "content": f"Subject: {subject}. Number of paragraphs: {number_of_paragpraphs}. Language: EN"
            }
        ],
        model="llama3-8b-8192",
    )

    response = chat_completion.choices[0].message.content
    print(f"Created video script: {response}")
    return response


def generate_search_terms(subject, script, amount=5):
    print(f"Generating {amount} search terms for script")
    chat_completion = client.chat.completions.create(
        messages=[
            {
                "role": "system",
                "content": PRE_SEARCH_TERMS_CREATION_SCRIPT
            },
            {
                "role": "user",
                "content": f"Subject: {subject}, Amount: {amount}, Script: {script}"
            }
        ],
        model="llama3-8b-8192"
    )

    response = chat_completion.choices[0].message.content
    print(f"Created search terms: {response}")
    return response



