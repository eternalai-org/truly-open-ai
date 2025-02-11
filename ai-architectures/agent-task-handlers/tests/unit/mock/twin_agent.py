def mock_random_example_tweets(knowledge_base_id: str, **kwargs):
    if knowledge_base_id == "kn_1":
        return [
            "Mock example tweet 1",
            "Mock example tweet 2",
            "Mock example tweet 3",
        ]
    return []
