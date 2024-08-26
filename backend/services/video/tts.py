import uuid
import requests
import base64
from concurrent.futures import ThreadPoolExecutor

VOICE_LANGUAGE = "en_au_002"
TEXT_LENGTH_LIMIT = 300  # Set to 299 since the max allowed is 299
AUDIO_DIRECTORY = "backend/services/video/saved/voice"
TTS_API_URL = "https://countik.com/api/text/speech"


def convert_text_to_speech(text):
    text_parts = _split_text(text, TEXT_LENGTH_LIMIT)

    with ThreadPoolExecutor() as executor:
        audio_base64_parts = list(executor.map(_generate_audio, text_parts))

    concatenated_audio_base64 = "".join(audio_base64_parts)
    _save_to_file(concatenated_audio_base64)
    print("Audio file saved successfully.")


def _split_text(text, limit):
    sentences = text.split('. ')
    parts = []
    current_part = ""

    for sentence in sentences:
        if len(current_part) + len(sentence) + 1 <= limit:  # +1 for the space after a period
            if current_part:
                current_part += " "
            current_part += sentence
        else:
            parts.append(current_part)
            current_part = sentence

    if current_part:
        parts.append(current_part)

    return parts


def _generate_audio(text_part):
    response = requests.post(TTS_API_URL, headers={
        "Content-Type": "application/json"
    }, json={"text": text_part, "voice": VOICE_LANGUAGE})
    return response.json()["v_data"]


def _save_to_file(base64_audio):
    audio_bytes = base64.b64decode(base64_audio)
    save_path = f"{AUDIO_DIRECTORY}/{uuid.uuid4()}.mp3"

    with open(save_path, "wb") as file:
        file.write(audio_bytes)
