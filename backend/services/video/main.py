from backend.services.video.gpt import generate_video_script, generate_search_terms
from backend.services.video.pexel import get_video_infos, filter_and_get_video_links, download_videos_and_get_paths


def main():
    subject = "Amazon Forest"
    script = generate_video_script(subject)
    search_terms = generate_search_terms(subject, script)
    video_links = set()
    for term in search_terms:
        videos = get_video_infos(term)
        filtered_videos_links = filter_and_get_video_links(videos)
        video_links.update(filtered_videos_links)

    if not video_links:
        print("No videos found to download")
        return

    saved_video_paths = download_videos_and_get_paths(video_links)
    sentences = [sentence for sentence in script.split(". ") if sentence]


main()
