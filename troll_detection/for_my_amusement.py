from tracker import Detector, Score

msgs = ["Donde esta bibliotecha",
        "I wanna fuck you in the ass",
        "nice tits",
        "You stupid poo",
        "I'm gonna give your wife cancer",
        "you dirty curry man"]

for msg in msgs:
    d = Detector()
    d.request(msg, True, "conv1")
    score = d.get_rolling_score_as_score(Detector.generate_id("conv1", True))
    single_score = Score.get_single_score_metric(score)

    print("Message:\t", msg)
    print("Score:\t", score)
    print("Single score metric:\t", single_score)
    print("\n")
