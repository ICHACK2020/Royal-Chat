from googleapiclient import discovery

API_KEY = 'AIzaSyA3l1ih4OQ11VvzJzkXjdkjCkbUOQWbjDE'


class Score:

    def __init__(self, user_id):
        self._user_id = user_id
        self._recent_score = {}
        scoring_features = ['THREAT',
                            'INSULT',
                            'IDENTITY_ATTACK',
                            'PROFANITY']

        self._num_points = 0
        self._scores = {key: 0 for key in scoring_features}

    def update(self, scores):
        self._recent_score = scores
        self._num_points += 1

        for key, value in scores.items():
            new_rolling_score = ((self._scores[key] * (self._num_points - 1)) + value) / self._num_points
            self._scores[key] = new_rolling_score

    @classmethod
    def get_single_score_metric(cls, scores):
        return sum(scores.values()) / len(scores.values())

    @property
    def user_id(self):
        return self._user_id

    @property
    def recent_score(self):
        return self._recent_score

    @property
    def rolling_scores(self):
        return self._scores


class Detector:
    def __init__(self):
        self._db = {}

    def request(self, msg, uid, convid):
        if uid in self._db.keys():
            self._db[uid].update(Detector.get_scores(msg))
        else:
            self._db[uid] = Score(uid)
            self._db[uid].update(Detector.get_scores(msg))

    @classmethod
    def get_scores(cls, msg):
        service = discovery.build('commentanalyzer', 'v1alpha1', developerKey=API_KEY)
        analyze_request = {
            'comment': {'text': msg},
            'requestedAttributes': {'INSULT': {},
                                    'THREAT': {},
                                    'PROFANITY': {},
                                    'IDENTITY_ATTACK': {}}
        }
        response = service.comments().analyze(body=analyze_request).execute()

        return {k: feature_dict['summaryScore']['value'] for k, feature_dict in response['attributeScores'].items()}

    def get_recent_score(self, uid):
        return Score.get_single_score_metric(self._db[uid].recent_score)

    def get_rolling_score(self, uid):
        return Score.get_single_score_metric(self._db[uid].main_score)
