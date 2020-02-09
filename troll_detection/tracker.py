from googleapiclient import discovery

API_KEY = 'AIzaSyA3l1ih4OQ11VvzJzkXjdkjCkbUOQWbjDE'


class Score:
    _scoring_features = ['THREAT',
                         'INSULT',
                         'IDENTITY_ATTACK',
                         'PROFANITY']

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
        weightings = {
            'THREAT': 48,
            'INSULT': 10,
            'IDENTITY_ATTACK': 38,
            'PROFANITY': 4
        }

        # Very simplistic algorithm, could use some refinement
        score = 0
        for key, weight in weightings.items():
            score += weight * scores[key]

        score /= 100

        return score

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
                                    'IDENTITY_ATTACK': {}}
        }
        response = service.comments().analyze(body=analyze_request).execute()

        return {k: feature_dict['summaryScore']['value'] for k, feature_dict in response['attributeScores'].items()}

    def get_recent_score(self, id):
        return Score.get_single_score_metric(self._db[id].recent_score)

    def get_rolling_score(self, id):
        return Score.get_single_score_metric(self._db[id].rolling_scores)
