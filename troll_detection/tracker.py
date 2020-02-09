from googleapiclient import discovery
from googleapiclient.errors import HttpError

API_KEY = 'AIzaSyA3l1ih4OQ11VvzJzkXjdkjCkbUOQWbjDE'


class Score:
    _scoring_features = ['THREAT',
                         'INSULT',
                         'IDENTITY_ATTACK',
                         'PROFANITY',
                         'SEXUALLY_EXPLICIT']

    def __init__(self):
        self._recent_score = {}
        self._num_points = 0
        self._scores = {key: 0 for key in Score._scoring_features}

    def update(self, scores):
        self._recent_score = scores
        self._num_points += 1

        for key, value in scores.items():
            new_rolling_score = ((self._scores[key] * (self._num_points - 1)) + value) / self._num_points
            self._scores[key] = new_rolling_score

    @classmethod
    def get_single_score_metric(cls, scores):
        for key, value in scores.items():
            if value > 0.8 and key != 'PROFANITY':
                return value

        return sum(scores.values()) / len(scores)

    @property
    def recent_score(self):
        return self._recent_score

    @property
    def rolling_scores(self):
        return self._scores


class Detector:
    def __init__(self):
        self._db = {}

    def request(self, msg, uid, conv_id):
        u_conv_id = Detector.generate_id(conv_id, uid)
        if uid in self._db.keys():
            self._db[u_conv_id].update(Detector.get_scores(msg))
        else:
            self._db[u_conv_id] = Score()
            self._db[u_conv_id].update(Detector.get_scores(msg))

    @classmethod
    def generate_id(cls, conv_id, uid):
        return conv_id + str(uid)

    @classmethod
    def get_scores(cls, msg):
        service = discovery.build('commentanalyzer', 'v1alpha1', developerKey=API_KEY)
        analyze_request = {
            'comment': {'text': msg},
            'requestedAttributes': {'INSULT': {},
                                    'THREAT': {},
                                    'PROFANITY': {},
                                    'IDENTITY_ATTACK': {},
                                    'SEXUALLY_EXPLICIT': {}}
        }

        try:
            response = service.comments().analyze(body=analyze_request).execute()
        except HttpError:
            # If there is an error, for now we just return 0.5. This is because the API only supports english
            return {k: 0.5 for k in ['INSULT', 'THREAT', 'PROFANITY', 'IDENTITY_ATTACK', 'SEXUALLY_EXPLICIT']}

        return {k: feature_dict['summaryScore']['value'] for k, feature_dict in response['attributeScores'].items()}

    def get_rolling_score_as_score(self, id):
        return self._db[id].rolling_scores

    def get_recent_score(self, id):
        return Score.get_single_score_metric(self._db[id].recent_score)

    def get_rolling_score(self, id):
        return Score.get_single_score_metric(self._db[id].rolling_scores)
