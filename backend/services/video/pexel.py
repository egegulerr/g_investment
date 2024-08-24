import uuid

from dotenv import load_dotenv
import os
import requests

load_dotenv()
PEXELS_API_KEY = os.getenv("PEXELS_API_KEY")
DOWNLOAD_DIRECTORY = "./saved"
PER_PAGE_QUERY_LIMIT = 15
MINIMUM_VIDEO_DURATION = 10


def get_video_infos(term):
    headers = {
        "Authorization": PEXELS_API_KEY
    }

    url = f"https://api.pexels.com/videos/search?query={term}&per_page={PER_PAGE_QUERY_LIMIT}"
    response = requests.get(url, headers=headers)
    return response.json()


def filter_and_get_video_links(videos):
    filtered_videos = set()

    for i in range(PER_PAGE_QUERY_LIMIT):
        if videos["videos"][i]["duration"] < MINIMUM_VIDEO_DURATION:
            continue

        video_files = videos["videos"][i]["video_files"]
        valid_videos = [video for video in video_files if ".com/video-files" in video["link"]]

        if valid_videos:
            max_res_video = max(valid_videos, key=lambda video: video['width'] * video['height'])
            url = max_res_video['link']
            filtered_videos.append(url)
            print("Found high res video. Appending it")

    return filtered_videos


def download_videos_and_get_paths(links):
    paths = []
    for link in links:

        video_id = uuid.uuid4()
        save_path = f"{DOWNLOAD_DIRECTORY}/{video_id}.mp4"

        print(f"Downloading a video with name {video_id}")

        response = requests.get(link)
        with open(save_path, "wb") as f:
            f.write(response.content)
        paths.append(save_path)

    return paths
