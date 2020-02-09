from tracker import Detector, Score

msgs = ["I have never seen anything so despicable before",
        "If you walked into a bar you would hit your head",
        "why are you offended you're such a snowflake",
        "When you speak I want to blow your brains out",
        "wow. genius. fucking brilliant",
        "I hope your son gets cancer",
        "I'm going to trace your location and find you",
        "If you don't believe in God, he will punish you in hell forever",
        "We need to take the country forward. make britain great again. We don't need the EU.",
        "That's so gay",
        "Well no obviously not, I think it's a little more subtle than that"]

for msg in msgs:
    d = Detector()
    d.request(msg, True, "conv1")
    score = d.get_rolling_score_as_score(Detector.generate_id("conv1", True))
    single_score = Score.get_single_score_metric(score)

    print("Message:\t", msg)
    print("Score:\t", score)
    print("Single score metric:\t", single_score)
    print("\n")
