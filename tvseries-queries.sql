#Queries
SELECT Actors.Name FROM actors LEFT JOIN actors_has_episodes ON idActor = Actors_idActor LEFT JOIN episodes ON idEpisodes = Episodes_idEpisodes LEFT JOIN tv_series ON idTV_Series = idSerie WHERE tv_series.Name = "Big Sister"


SELECT directors.Name, COUNT(directors.idDirectors) AS 'Count' FROM directors LEFT JOIN directors_has_episodes ON idDirectors = Directors_idDirectors Group BY idDirectors ORDER BY Count DESC LIMIT 1
