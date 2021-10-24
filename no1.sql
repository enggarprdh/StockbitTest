SELECT u1.ID, u1.UserName, u2.UserName parentUsername
FROM USER u1 
LEFT JOIN USER u2 ON u2.ID = u1.Parent

