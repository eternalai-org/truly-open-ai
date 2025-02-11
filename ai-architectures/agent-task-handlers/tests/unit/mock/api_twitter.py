from x_content.wrappers.api.twitter_v2.models.response import GenerateActionDto, Response


def mock_tweet(**kwargs):
    return Response(data=GenerateActionDto(success=True))
