from tracker import Detector, Score

msgs = ["Donde esta bibliotecha"]

for msg in msgs:
    d = Detector()
    d.request(msg, True, "conv1")
    score = d.get_rolling_score_as_score(Detector.generate_id("conv1", True))
    single_score = Score.get_single_score_metric(score)

    print("Message:\t", msg)
    print("Score:\t", score)
    print("Single score metric:\t", single_score)
    print("\n")
